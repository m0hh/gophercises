package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type link struct {
	url  string
	text string
}

func main() {
	file, err := os.Open("ex1.html")

	if err != nil {
		panic("cannot open file")
	}

	var links []link
	htmlll := html.NewTokenizer(file)

	for {
		tt := htmlll.Next()
		if tt == html.ErrorToken {
			break
		}

		data := htmlll.Token().Data

		if htmlll.Token().Type.String() == "StartTag" && data == "a" {
			tt := htmlll.Next()
			if tt == html.ErrorToken {
				break
			}

			links = append(links, link{url: htmlll.Token().Data})
		}
	}
	fmt.Println(links)
}
