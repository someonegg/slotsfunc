// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slotsfunc

import "testing"

func TestAllotUnion(t *testing.T) {
	first := Allot(nil, []Slot{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, []Inst{"a", "b", "c"}, nil)
	t.Log(first)

	second := Allot(first, nil, []Inst{"d", "e"}, nil)
	t.Log(second)

	third := Allot(second, nil, nil, []Inst{"a", "e", "f"})
	t.Log(third)

	fourth := Allot(third, []Slot{1, 2, 3, 4, 5}, nil, nil)
	t.Log(fourth)

	fifth := Allot(fourth, nil, []Inst{"g", "h"}, []Inst{"d"})
	t.Log(fifth)

	sixth := Allot(fifth, []Slot{4, 3, 2, 5, 1}, nil, nil)
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
}
