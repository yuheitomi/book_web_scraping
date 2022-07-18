package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly/debug"
)

func postGet() {
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	cookies := login()

	fmt.Println("Start scraping after login page")

	c := colly.NewCollector()
	if err := c.SetCookies("https://pythonscraping.com/pages/cookies/welcome.php", cookies); err != nil {
		log.Fatal(err)
	}

	c.OnResponse(func(resp *colly.Response) {
		log.Println("Response received ", resp.StatusCode)
		log.Println(resp.Request.URL)
		log.Println(string(resp.Body))
	})

	if err := c.Visit("http://pythonscraping.com/pages/cookies/profile.php"); err != nil {
		log.Fatal(err)
	}
}

func login() []*http.Cookie {
	c := colly.NewCollector()

	auth := map[string]string{
		"username": "Yuhei",
		"password": "password",
	}
	c.OnResponse(func(resp *colly.Response) {
		log.Println("Response received ", resp.StatusCode)
		log.Println(resp.Request.URL)
		log.Println(string(resp.Body))
	})
	err := c.Post("https://pythonscraping.com/pages/cookies/welcome.php", auth)
	if err != nil {
		log.Fatal(err)
	}

	cookies := c.Cookies("http://pythonscraping.com/pages/cookies")
	return cookies
}
