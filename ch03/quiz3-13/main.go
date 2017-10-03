package main

import (
	"fmt"
)

const (
	KiB = 1000
	MiB = KiB * KiB
	GiB = MiB * KiB
	TiB = GiB * KiB
	PiB = TiB * KiB
	EiB = PiB * KiB
	ZiB = EiB * KiB
	YiB = ZiB * KiB
)

func main() {
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB) //, ZB, YB)
}
