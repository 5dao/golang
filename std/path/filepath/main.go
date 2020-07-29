package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	arg0 := os.Args[0]
	log.Println(arg0)
	log.Println((filepath.Base(arg0)))

	log.Println(filepath.Ext(arg0))
}
