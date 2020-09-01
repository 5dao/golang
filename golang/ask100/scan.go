package ask100

import (
	"fmt"
)

// func main() {
// 	var a interface{}
// 	fmt.Scan(&a)
// 	fmt.Println(a)
// }

// Data Data
type Data struct {
	ID int
}

// func main() {
// 	var a Data
// 	fmt.Scan(&a)
// 	fmt.Println(a)
// }

// Scan Scan
func Scan() {
	var a string
	fmt.Scan(&a) // if not letter ,only space or return key ???
	fmt.Println(a)
}
