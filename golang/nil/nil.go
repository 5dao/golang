//Package nil nil
package nil

import (
	"fmt"
	"unsafe"
)

// Da da
type Da struct {
}

// T0 t0
func T0() {
	d := (*int)(nil)
	fmt.Println(d == nil)
}

// T1 T1
func T1() {
	var t *Da
	var i interface{} = t

	t2 := ToIS(t)
	fmt.Println(t2.pt, t2.pv)

	i2 := ToIS(i)
	fmt.Println(i2.pt, i2.pv)

	fmt.Println(t == nil, i == t, i == nil)
}

// T2 T2
func T2() {
	var t interface{}
	t1 := ToIS(t)
	fmt.Println(t1.pt, t1.pv)
	fmt.Println("-------", t == nil)

	t = 1
	t2 := ToIS(t)
	fmt.Println(t2.pt, t2.pv)
	fmt.Println("-------", t == nil)

	var d *Da
	d2 := ToIS(d)
	fmt.Println(d2.pt, d2.pv)
	fmt.Println("-------", d == nil)
}

// IS IS
type IS struct {
	pt uintptr
	pv uintptr
}

// ToIS ToIS
func ToIS(i interface{}) IS {
	return *(*IS)(unsafe.Pointer(&i))
}
