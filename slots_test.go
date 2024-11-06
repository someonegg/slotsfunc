// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slotsfunc

import "testing"

type Slot = int
type Inst = string

func TestAllotUnion(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		first := Allot(nil, []Slot{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}, []Inst{"a", "b"}, nil)
		t.Log(first)

		second := Allot(first, nil, []Inst{"c", "d", "e"}, nil)
		t.Log(second)

		third := Allot(second, nil, nil, []Inst{"b", "e", "f"})
		t.Log(third)

		fourth := Allot(third, []Slot{1, 2, 3, 4, 5}, nil, nil)
		t.Log(fourth)

		fifth := Allot(fourth, nil, []Inst{"g"}, []Inst{"d"})
		t.Log(fifth)

		sixth := Allot(fifth, []Slot{4, 3, 2, 5, 1}, []Inst{"h", "i"}, nil)
		t.Log(sixth)

		unionth := Union(fifth, sixth)
		t.Log(unionth)

		final := Reverse(unionth)
		t.Log(final)

		replicas := 0
		for _, insts := range final {
			if replicas == 0 {
				replicas = len(insts)
			}
			if len(insts) != replicas {
				t.Fail()
			}
		}
	})

	t.Run("case 2", func(t *testing.T) {
		first := Allot(nil, []Slot{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []Inst{"a", "b", "c"}, nil)
		t.Log(first)

		second := Allot(first, nil, []Inst{"d"}, nil)
		t.Log(second)

		third := Allot(second, nil, []Inst{"e", "f", "g"}, []Inst{"a"})
		t.Log(third)
	})

	t.Run("case 3", func(t *testing.T) {
		first := Allot(nil, []Slot{1, 2, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}, []Inst{"a", "b", "c"}, nil)
		t.Log(first)

		second := Allot(first, nil, []Inst{"d"}, nil)
		t.Log(second)

		third := Allot(second, nil, []Inst{"e", "f", "g"}, []Inst{"a"})
		t.Log(third)
	})

	t.Run("case 4", func(t *testing.T) {
		first := Allot(nil, []Slot{1, 2}, []Inst{"a", "b", "c"}, nil)
		t.Log(first)

		second := Allot(first, nil, []Inst{"d"}, nil)
		t.Log(second)

		third := Allot(second, nil, []Inst{"e", "f", "g"}, []Inst{"a"})
		t.Log(third)
	})
}
