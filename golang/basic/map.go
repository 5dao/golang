package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	mapC()

// 	g()
// }

func mapC() {
	defer func() {
		if rev := recover(); rev != nil {
			fmt.Println("mapc", rev)
		}
	}()

	*(*int)(nil) = 0
}

func g() {
	defer func() {
		if rev := recover(); rev != nil {
			fmt.Println(rev)
		}
	}()
	m := map[int]interface{}{}

	wn := 4

	var wg sync.WaitGroup
	wg.Add(wn)

	for i := 0; i < wn; i++ {
		go func() {
			j := 0
			for {
				m[j] = j
				j++
				if j > 100000 {
					break
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

}
