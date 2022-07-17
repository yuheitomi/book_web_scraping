package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

func getExternalLinks(r io.Reader, link *url.URL) []string {
	var result []string
	// TODO: implement here
	return result
}

func getInternalLinks(r io.Reader, link *url.URL) []string {
	var result []string
	// TODO: implement here
	return result
}

func getRandomExternalLink(link *url.URL) string {
	resp, err := http.Get(link.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	externalLinks := getExternalLinks(resp.Body, link)
	if len(externalLinks) == 0 {
		fmt.Println("No external link found.")
		domain := fmt.Sprintf("%v://%v", link.Scheme, link.Host)
		internalLinks := getInternalLinks(resp.Body, domain)
		return getRandomExternalLink(internalLinks[rand.Intn(len(internalLinks))])
	} else {
		return externalLinks[rand.Intn(len(externalLinks))]
	}
}

func followExternalLinkOnly(ref string) {
	startUrl, err := url.Parse(ref)
	if err != nil {
		log.Fatal(err)
	}
	externalLink := getRandomExternalLink(startUrl)
	fmt.Printf("Random external link: %v", externalLink)
	followExternalLinkOnly(externalLink)
}
