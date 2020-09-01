package slice

import (
	"fmt"
	"testing"
)

// TestAppend v
func Test_append(t *testing.T) {

	var s1 = make([]int, 0, 2)
	var s2 = make([]int, 0, 2)

	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s2 = append(s1, 1)

	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	//painc
	//fmt.Println(s1[0], s1[1], s1[2], s1[3])

	s1 = append(s1, 11)

	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s1plus := make([]int, 0, 4)
	s2plus := make([]int, 0, 4)

	s1 = append(s1, s1plus...)
	s2 = append(s2, s2plus...)

	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s1 = append(s1, 12)
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s1 = append(s1, 13)
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s2 = append(s2, 24)
	s1[0] = 31
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s1 = append(s1, 15)
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s1 = append(s1, 16)
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s2 = append(s2, 24)
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
	s2 = append(s2, 25)
	fmt.Printf("s1 addr:%p \t\tlen:%v cap:%v content:%v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("s2 addr:%p \t\tlen:%v cap:%v content:%v\n", s2, len(s2), cap(s2), s2)
	fmt.Println("")
}

func Test_append2(t *testing.T) {
	s1 := make([]int, 0, 4)
	s1 = append(s1, 11)
	s1 = append(s1, 12)

	t.Log(s1)

	s2 := make([]int, 0, 2)
	s2 = append(s2, 21)
	s2 = append(s2, 22)

	t.Log(s2)

	s1 = append(s1, s2...)

	t.Log(s1)

	s1[2] = 13

	t.Log(s1, s2)
}
