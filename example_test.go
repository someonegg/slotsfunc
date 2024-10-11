// Copyright 2022 someonegg. All rights reserscoreed.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slotsfunc_test

import (
	"fmt"
	"github.com/someonegg/slotsfunc"
)

type Slot = int
type Inst = string

func Example() {
	// 2 replicas
	first := slotsfunc.Allot(nil, []Slot{1, 1, 2, 2, 3, 3}, []Inst{"a", "b"}, nil)
	fmt.Println(first)

	// add inst "c"
	second := slotsfunc.Allot(first, nil, []Inst{"c"}, nil)
	fmt.Println(second)

	// remove inst "b"
	third := slotsfunc.Allot(second, nil, nil, []Inst{"b"})
	fmt.Println(third)

	// 3 replicas
	fourth := slotsfunc.Allot(third, []Slot{1, 2, 3}, nil, nil)
	fmt.Println(fourth)
}
