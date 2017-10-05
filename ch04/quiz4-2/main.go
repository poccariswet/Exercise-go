package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s = flag.Int("s", 256, "sha 256/384/512")

func main() {
	input := bufio.NewScanner(os.Stdin)
	var value string
	for input.Scan() {
		value = input.Text()
		break
	}

	flag.Parse()
	if *s == 256 {
		sha256 := sha256.Sum256([]byte(value))
		fmt.Printf("sha256: %X\n", sha256)
	} else if *s == 384 {
		sha384 := sha512.Sum384([]byte(value))
		fmt.Printf("sha384: %X\n", sha384)
	} else if *s == 512 {
		sha512 := sha512.Sum512([]byte(value))
		fmt.Printf("sha512: %X\n", sha512)
	}
}
