package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	var b bytes.Buffer

	for n > 0 {
		l := n % 3
		if l == 0 {
			l = 3
		}
		b.WriteString(s[:l])
		if n > 3 {
			b.WriteString(",")
		}
		s = s[l:]
		n = len(s)
	}

	return b.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}
