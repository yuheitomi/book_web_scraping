package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func getExternalLinks(doc *goquery.Document, excludeHost string) []string {
	uniqueLinks := make(map[string]bool)
	excludeM := regexp.MustCompile(fmt.Sprintf("^(http|www).+(%v).*$", excludeHost))
	includeM := regexp.MustCompile("^(http|www).+$")

	doc.Find("a").Each(func(i int, sel *goquery.Selection) {
		href := sel.AttrOr("href", "")
		if excludeM.Match([]byte(href)) {
			return
		}
		if !includeM.Match([]byte(href)) {
			return
		}

		_, ok := uniqueLinks[href]
		if !ok {
			uniqueLinks[href] = true
		}
	})

	result := make([]string, 0, len(uniqueLinks))
	for k := range uniqueLinks {
		result = append(result, k)
	}
	return result
}

func getInternalLinks(r io.Reader, link string) []string {
	var result []string
	// TODO: implement here
	return result
}

func getRandomExternalLinks(link string) []string {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	linkUrl, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	externalLinks := getExternalLinks(doc, linkUrl.Host)
	if len(externalLinks) == 0 {
		fmt.Println("No external link found.")
		linkUrl, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}

		domain := fmt.Sprintf("%v://%v", linkUrl.Scheme, linkUrl.Host)
		internalLinks := getInternalLinks(resp.Body, domain)
		if len(internalLinks) > 0 {
			return getRandomExternalLinks(internalLinks[rand.Intn(len(internalLinks))])
		} else {
			return nil
		}
	} else {
		return externalLinks
	}
}

func followExternalLinkOnly(ref string) {
	for {
		externalLinks := getRandomExternalLinks(ref)
		if externalLinks == nil {
			break
		}
		i := rand.Intn(len(externalLinks))
		ref = externalLinks[i]
		fmt.Printf("Random external link (%d/%d): %v\n", i, len(externalLinks), ref)
	}
}
