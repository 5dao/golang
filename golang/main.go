package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	//838392392929329392392392939292392392932823238238283283283823828328
	bg, ok := new(big.Int).SetString("838392392929329392392392939292392392932823238238283283283823828328", 10)

	fmt.Println(bg, ok)
	fmt.Println(bg.String())

	return

	addrLen := 1
	getAddrPercent := 23

	for i := 1; i < 25; i++ {
		nu := i * getAddrPercent / 100
		fmt.Println(i, nu)
	}

	numAddresses := math.Ceil(float64(addrLen*getAddrPercent) / 100)

	fmt.Println(numAddresses)
}
