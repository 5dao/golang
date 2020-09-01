package main

import (
	"fmt"

	"golang.org/x/crypto/acme"
)

const (
	contactEmail = "i.5d@qq.com"
	domain       = "dao.bookplex.cn"
)

func main() {
	c := acme.Client{}

	fmt.Println(c)
}
