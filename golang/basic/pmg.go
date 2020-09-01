 package main 


// import(
// 	"runtime"
// 	"sync"
// 	"fmt"
	
// )
// func main() {
// 	runtime.GOMAXPROCS(1)

// 	wg:=sync.WaitGroup{}



// 	wg.Add(30)
// 	fmt.Println("ok1")
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			fmt.Println("ii: ", i)
// 			wg.Done()
// 		}()
// 	}
// 	fmt.Println("ok2")
// 	for i := 0; i < 20; i++ {
// 		// fmt.Println("ok2",i)
// 		go func(i int) {
// 			fmt.Println("i: ", i)
// 			wg.Done()
// 		}(i)
// 		// fmt.Println("ok2..",i)
// 	}
// 	fmt.Println("ok3")
// 	 wg.Wait()

	
// }