// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slotsfunc provides several functions to manage slots.
package slotsfunc

import (
	"math"
	"sort"
)

type Table[Inst, Slot comparable] map[Inst][]Slot

type instSlotLen[Inst comparable] struct {
	Inst    Inst
	SlotLen int
}

func Allot[Inst, Slot comparable](base Table[Inst, Slot], added []Slot, news []Inst, rms []Inst) Table[Inst, Slot] {
	t := make(Table[Inst, Slot])

	total := len(added)
	for inst, slots := range base {
		t[inst] = append([]Slot{}, slots...)
		total += len(slots)
	}

	allots := append([]Slot{}, added...)
	for _, inst := range rms {
		slots, ok := t[inst]
		if !ok {
			continue
		}
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
	avgCeil := int(math.Ceil(float64(total) / float64(len(t))))

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

	sortedInst := make([]instSlotLen[Inst], 0, len(t))
	for inst, slots := range t {
		for len(slots) > avgCeil {
			evictRepeated(slots)
			allots = append(allots, slots[len(slots)-1])
			slots = slots[0 : len(slots)-1]
		}
		t[inst] = slots
		sortedInst = append(sortedInst, instSlotLen[Inst]{inst, len(slots)})
	}
	sort.Slice(sortedInst, func(i, j int) bool {
		return sortedInst[i].SlotLen > sortedInst[j].SlotLen
	})
	for need := len(news)*avg - len(allots); need > 0; {
		noop := true
		for _, instLen := range sortedInst {
			inst := instLen.Inst
			slots := t[inst]
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

func Union[Inst, Slot comparable](a, b Table[Inst, Slot]) Table[Inst, Slot] {
	t := make(Table[Inst, Slot])
	for inst, slots := range a {
		t[inst] = append([]Slot{}, slots...)
	}
	for inst, slots := range b {
		t[inst] = append(t[inst], slots...)
	}
	return t
}

func Reverse[Inst, Slot comparable](t Table[Inst, Slot]) Table[Slot, Inst] {
	r := make(Table[Slot, Inst])
	for inst, slots := range t {
		for _, slot := range slots {
			r[slot] = append(r[slot], inst)
		}
	}
	return r
}
