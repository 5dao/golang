//Package main
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//main main
func main() {

	var a *Da = &Da{}

	pa := unsafe.Pointer(a)

	fmt.Println(reflect.ValueOf(a))
	fmt.Println(pa)

}

type Da struct {
}
