// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slotsfunc

import "testing"

func TestAllotUnion(t *testing.T) {
	first := Allot(nil, []Slot{1, 1, 2, 2, 3, 3}, []Inst{"a", "b", "c"}, nil)
	t.Log(first)

	second := Allot(first, nil, []Inst{"d", "e"}, nil)
	t.Log(second)

	third := Allot(second, nil, nil, []Inst{"a", "e", "f"})
	t.Log(third)

	fourth := Allot(third, []Slot{1, 2, 3}, nil, nil)
	t.Log(fourth)

	fifth := Allot(fourth, nil, []Inst{"g", "h"}, []Inst{"d"})
	t.Log(fifth)

	sixth := Allot(fifth, []Slot{3, 2, 1}, nil, nil)
	t.Log(sixth)

	final := Union(fifth, sixth)
	t.Log(final)
}
