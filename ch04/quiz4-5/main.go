package main

import "fmt"

func depremove(arr []string) []string {
	i := 0
	prev := ""
	for _, s := range arr {
		if s != prev {
			arr[i] = s
			i++
			prev = s
		}
	}
	return arr[:i]
}

func main() {
	var arr = []string{"hello", "world", "hello", "hello", "hello"}
	res := depremove(arr)
	fmt.Println(res)
}
