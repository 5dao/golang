package pointer

import (
	"log"
	"testing"
)

type Foo struct {
	Name string
}

func change(foo Foo) Foo {
	log.Printf("change %p,%s", &foo, foo.Name)
	foo.Name = "changed"
	return foo
}

func changePointer(foo *Foo) Foo {
	log.Printf("changePointer %p,%s", foo, foo.Name)
	foo.Name = "changed"
	log.Printf("changePointer %p,%s", foo, foo.Name)
	return *foo
}

func Test_m(t *testing.T) {
	foo := Foo{Name: "test"}
	log.Printf("%p,%s", &foo, foo.Name)

	// fooo := change(foo)
	// log.Printf("%p,%s", &fooo, fooo.Name)
	// log.Printf("%p,%s", &foo, foo.Name)

	fooo := changePointer(&foo)
	log.Printf("%p,%s", &fooo, fooo.Name)
	fooo.Name = "final"
	log.Printf("%p,%s", &fooo, fooo.Name)
	log.Printf("%p,%s", &foo, foo.Name)
}
