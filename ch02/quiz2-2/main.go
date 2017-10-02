// 数値引数を摂氏と華氏へ変換する
package main

import (
	"fmt"
	"os"
	"strconv"

	"./tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		// °F = °C
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		// ft = m
		ft := tempconv.Feet(t)
		m := tempconv.Meter(t)

		// pound = kg
		pd := tempconv.Pound(t)
		kg := tempconv.Kilogram(t)

		fmt.Println("°F <=> °C")
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Println("ft <=> m")
		fmt.Printf("%s = %s, %s = %s\n", ft, tempconv.FtToM(ft), m, tempconv.MToFt(m))
		fmt.Println("pound <=> kg")
		fmt.Printf("%s = %s, %s = %s\n", pd, tempconv.PdToKg(pd), kg, tempconv.KgToPd(kg))
	}
}
