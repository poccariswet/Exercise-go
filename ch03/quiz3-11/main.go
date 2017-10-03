package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	fl := ""
	num := strings.Index(s, ".")
	if num > 0 {
		fl = s[num:]
		s = s[:num]
	}

	n := len(s)
	for n > 0 {
		l := n % 3
		if l == 0 {
			l = 3
		}
		buf.WriteString(s[:l])
		if n > 3 {
			buf.WriteString(",")
		}
		s = s[l:]
		n = len(s)
	}
	buf.WriteString(fl)

	return buf.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}
