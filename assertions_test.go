package utils

import (
	"testing"
)

// Ref: gofiber/utils
func TestAssertEqual(t *testing.T) {
	t.Parallel()
	AssertEqual(nil, []string{}, []string{})
	AssertEqual(t, []string{}, []string{})
}

func TestAssertPanics(t *testing.T) {
	t.Parallel()
	var a []int
	AssertPanics(nil, "should panic when index out of range", func() {
		_ = a[1]
	})
	AssertPanics(t, "should panic when index out of range", func() {
		_ = a[1]
	})
}

type tester interface {
	do()
}

func TestIsNil(t *testing.T) {
	t.Parallel()
	var map1 map[int]bool
	AssertEqual(t, true, IsNil(map1))
	map2 := make(map[int]bool)
	AssertEqual(t, false, IsNil(map2))

	var ch1 chan int
	AssertEqual(t, true, IsNil(ch1))
	ch2 := make(chan struct{})
	AssertEqual(t, false, IsNil(ch2))

	var slice1 []string
	AssertEqual(t, true, IsNil(slice1))
	slice2 := slice1
	AssertEqual(t, true, IsNil(slice2))
	slice1 = append(slice1, "")
	AssertEqual(t, false, IsNil(slice1))
	AssertEqual(t, true, IsNil(slice2))

	slice3 := make([]int, 0, 0)
	AssertEqual(t, false, IsNil(slice3))
	slice4 := slice3
	AssertEqual(t, false, IsNil(slice4))
	slice3 = nil
	AssertEqual(t, true, IsNil(slice3))
	AssertEqual(t, false, IsNil(slice4))

	var iface1 interface{}
	AssertEqual(t, true, IsNil(iface1))
	iface1 = nil
	AssertEqual(t, true, IsNil(iface1))
	iface1 = map1
	AssertEqual(t, true, IsNil(iface1))
	iface1 = map2
	AssertEqual(t, false, IsNil(iface1))
	var iface2 interface{} = (*int)(nil)
	AssertEqual(t, true, IsNil(iface2))

	var eface1 tester
	AssertEqual(t, true, IsNil(eface1))
	var eface2 = new(tester)
	AssertEqual(t, false, IsNil(eface2))

	var ptr *int
	AssertEqual(t, true, IsNil(ptr))

	var err error
	AssertEqual(t, true, IsNil(err))

	var fun func(int) error
	AssertEqual(t, true, IsNil(fun))

	var struct1 *TChoice
	AssertEqual(t, true, IsNil(struct1))
	var struct2 TChoice
	AssertEqual(t, false, IsNil(struct2))
	var struct3 = &TChoice{}
	AssertEqual(t, false, IsNil(struct3))

	var s string
	AssertEqual(t, false, IsNil(s))
}
