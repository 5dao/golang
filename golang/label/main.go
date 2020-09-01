//Package l
package main

import "fmt"

//main main
func main() {
	do()
}

//do do
func do() {
	i := 1
	fmt.Println(i)

	if i > 2 {
		goto RS
	}

RS:

	fmt.Println("ok")
}
