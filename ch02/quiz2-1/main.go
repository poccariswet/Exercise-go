package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	c := tempconv.FToC(200)
	f := tempconv.KToF(200)
	k := tempconv.CToK(200)
	fmt.Printf("c : %v \nf : %v \nk : %v\n", c.String(), f.String(), k.String())
}
