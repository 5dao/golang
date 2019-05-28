package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	var listen string
	var dir string

	flag.StringVar(&listen, "l", ":81", "-l 0.0.0.0:81")
	flag.StringVar(&dir, "d", ".", "-d path/to/")
	flag.Parse()

	err := http.ListenAndServe(listen, http.FileServer(http.Dir(dir)))
	if err != nil {
		fmt.Println(err)
	}
}
