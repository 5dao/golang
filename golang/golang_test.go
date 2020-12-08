package golang

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {

	C()
}

func C() {
	fmt.Println("abc")

EXIT:
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	ch := make(chan int)
	<-ch
	goto EXIT
}

func TestVar(t *testing.T) {

	VarMain()

}
