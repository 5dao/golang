package main

import("time"
"fmt")

func main(){


	var c chan int

	go func(){
		 fmt.Println("a")
		c<-1
		fmt.Println("b")
	}()
	

time.Sleep(time.Second*10)
}

