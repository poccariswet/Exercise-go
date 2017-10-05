package main

import "fmt"

func main() {
	s := []byte{1, 2, 3, 4, 5}
	fmt.Println(reverse(s))
}

func reverse(bytes []byte) []byte {
	runes := []rune(string(bytes))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	copy(bytes, []byte(string(runes)))

	return bytes
}
