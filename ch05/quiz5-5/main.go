package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		word, img, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("CountWordsAndImages error:%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("word: %v\n", word)
		fmt.Printf("image: %v\n", img)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parse HTML error: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		in := bufio.NewScanner(strings.NewReader(n.Data))
		in.Split(GetWords)
		for in.Scan() {
			words++
		}
	}
	wd, im := countWordsAndImages(n.FirstChild)
	words += wd
	images += im
	wd, im = countWordsAndImages(n.NextSibling)
	words += wd
	images += im
	return
}

func GetWords(data []byte, non bool) (get int, non2 []byte, err error) {
	first := 0
	for width := 0; first < len(data); first += width {
		var r rune
		r, width = utf8.DecodeRune(data[first:])
		if unicode.IsLetter(r) {
			break
		}
	}
	for width, j := 0, first; first < len(data); j += width {
		var r rune
		r, width = utf8.DecodeRune(data[j:])
		if !unicode.IsLetter(r) {
			return j + width, data[first:j], nil
		}
	}
	if non && len(data) > first {
		return len(data), data[first:], nil
	}
	return first, nil, nil
}
