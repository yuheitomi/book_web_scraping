package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	executeMain()
}

func executeMain() {
	resp, err := http.Get("https://en.wikipedia.org/wiki/Kevin_Bacon")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	urls := extractLinks(resp.Body)

	for _, u := range urls {
		fmt.Println(u)
	}
}

func extractLinks(r io.Reader) (result []string) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div #bodyContent a").
		Each(func(i int, sel *goquery.Selection) {
			link, ok := sel.Attr("href")
			if !ok || !strings.Contains(link, "/wiki") {
				return
			}
			result = append(result, link)
		})
	return result
}
