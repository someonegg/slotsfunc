// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slotsfunc provides several functions to manage slots.
package slotsfunc

import "math"

type Inst interface{}
type Slot interface{}

type Table map[Inst][]Slot

func Allot(base Table, added []Slot, news []Inst, rms []Inst) Table {
	t := make(Table)

	total := len(added)
	for inst, slots := range base {
		t[inst] = append([]Slot{}, slots...)
		total += len(slots)
	}

	allots := append([]Slot{}, added...)
	for _, inst := range rms {
		slots := t[inst]
		allots = append(allots, slots...)
		delete(t, inst)
	}

	for _, inst := range news {
		t[inst] = []Slot{}
	}

	if total <= 0 || len(t) <= 0 {
		return t
	}

	avg := int(math.Floor(float64(total) / float64(len(t))))
	if avg < 1 {
		avg = 1
	}

	hasSlot := func(ss []Slot, s Slot) bool {
		for i := 0; i < len(ss); i++ {
			if ss[i] == s {
				return true
			}
		}
		return false
	}

	evictRepeated := func(ss []Slot) {
		for i := 0; i < len(ss); i++ {
			if hasSlot(ss[i+1:], ss[i]) {
				ss[i], ss[len(ss)-1] = ss[len(ss)-1], ss[i]
				return
			}
		}
	}

	for need := len(news)*avg - len(allots); need > 0; {
		noop := true
		for inst, slots := range t {
			if len(slots) > avg {
				evictRepeated(slots)
				allots = append(allots, slots[len(slots)-1])
				t[inst] = slots[0 : len(slots)-1]
				noop = false
				if need--; need <= 0 {
					break
				}
			}
		}
		if noop {
			break
		}
	}

	assign := func(new Slot, filter func(inst Inst, slots []Slot, new Slot) bool) bool {
		for inst, slots := range t {
			if filter(inst, slots, new) {
				t[inst] = append(slots, new)
				return true
			}
		}
		return false
	}

	for _, allot := range allots {
		if assign(allot, func(inst Inst, slots []Slot, new Slot) bool {
			return len(slots) < avg && !hasSlot(slots, new)
		}) {
			continue
		}

		if assign(allot, func(inst Inst, slots []Slot, new Slot) bool {
			return len(slots) == avg && !hasSlot(slots, new)
		}) {
			continue
		}

		if assign(allot, func(inst Inst, slots []Slot, new Slot) bool {
			return len(slots) < avg
		}) {
			continue
		}

		if !assign(allot, func(inst Inst, slots []Slot, new Slot) bool {
			return len(slots) == avg
		}) {
			panic("impossible")
		}
	}

	return t
}

func Union(a, b Table) Table {
	t := make(Table)
	for inst, slots := range a {
		t[inst] = append([]Slot{}, slots...)
	}
	for inst, slots := range b {
		t[inst] = append(t[inst], slots...)
	}
	return t
}

type Rtable map[Slot][]Inst

func Reverse(t Table) Rtable {
	r := make(Rtable)
	for inst, slots := range t {
		for _, slot := range slots {
			r[slot] = append(r[slot], inst)
		}
	}
	return r
}
