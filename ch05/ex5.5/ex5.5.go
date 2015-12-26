package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					images++
				}
			}
		}
	} else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		// Set the split function for the scanning operation.
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childwords, childimages := countWordsAndImages(c)
		words += childwords
		images += childimages
	}
	return words, images
}

//go run ex5.5/ex5.5.go https://news.ycombinator.com/item?id=10781548
func main() {
	url := "https://golang.org"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	fmt.Println("Parsing", url)
	words, images, _ := CountWordsAndImages(url)
	fmt.Println("Words:", words, "images:", images)
}
