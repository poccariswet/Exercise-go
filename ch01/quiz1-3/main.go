package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo_range() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo_strings() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func measurement(function_time func()) {
	start := time.Now()
	function_time()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func main() {
	measurement(echo_range)
	measurement(echo_strings)
}
