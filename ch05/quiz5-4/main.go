package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	data, err := GatherLink(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("%v\n", v)
	}
}

func visit(n *html.Node, result []string) []string {
	if n == nil {
		return result
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "src" {
				result = append(result, a.Val)
			}
		}
	}
	result = visit(n.FirstChild, result)
	result = visit(n.NextSibling, result)
	return result
}

func GatherLink(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var result []string
	result = visit(doc, result)
	return result, nil
}
