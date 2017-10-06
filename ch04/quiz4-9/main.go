package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countWord := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		countWord[input.Text()]++
	}
	for word, count := range countWord {
		fmt.Printf("%s : %dコ\n", word, count)
	}
}
