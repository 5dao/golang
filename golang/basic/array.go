package main

// func main() {

// 	// nums := [2]int{1, 2}
// 	// a(nums)
// 	// fmt.Println(nums)

// 	// nums2 := []int{1, 2}
// 	// b(nums2)
// 	// fmt.Println(nums2)

// 	p1 := People{Name: "p1"}
// 	c(p1)
// 	fmt.Println(p1.Name)

// 	p2 := &People{Name: "p2"}
// 	d(p2)
// 	fmt.Println(p2.Name)
// }

// array
func a(a [2]int) {
	a[0] = 100
	a[1] = 200
}

// slice
func b(a []int) {
	a[0] = 100
	a[1] = 200
}

//People People
type People struct {
	Name string
}

func c(p People) {
	p.Name = "cc"
}

func d(p *People) {
	p.Name = "cc"
}
