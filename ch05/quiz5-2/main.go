package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	data, err := Count(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("map----------------------------------")
	for i, d := range data {
		fmt.Printf("%v: %v\n", i, d)
	}
}

func visit(n *html.Node, result map[string]int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		result[n.Data]++
	}
	visit(n.FirstChild, result)
	visit(n.NextSibling, result)
}

func Count(r io.Reader) (map[string]int, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int)
	visit(doc, result)
	return result, nil
}
