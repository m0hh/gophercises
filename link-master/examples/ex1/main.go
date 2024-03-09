package main

import (
	"fmt"
	"link-parsing-html/link"
	"os"
)

func main() {
	// r := strings.NewReader(exampleHtml)
	r, err := os.Open("ex3.html")
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
