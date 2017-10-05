package main

import (
	"unicode"
	"unicode/utf8"
)

func condense(bytes []byte) []byte {
	var (
		ru   rune
		size int
	)
	ascii := byte(0x20)
	count := 0
	for i := 0; i < len(bytes); {
		ru, size = utf8.DecodeRune(bytes[i:])
		if unicode.IsSpace(ru) {
			bytes[count] = ascii
			count++
			i += size

			// 残りの空白は捨てる
			for i < len(bytes) {
				ru, size = utf8.DecodeRune(bytes[i:])
				if !unicode.IsSpace(ru) {
					break
				}
				i += size
			}
		} else {
			for j := 0; j < size; j++ {
				bytes[count] = bytes[i]
				count++
				i++
			}
		}
	}
	return bytes[:count]
}

func main() {
}
