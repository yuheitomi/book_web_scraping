package main

import "fmt"

var chapter = 10

func main() {
	switch chapter {
	case 3:
		followExternalLinkOnly("http://oreilly.com")
	case 10:
		postGet()
	default:
		fmt.Println("No chapter selected.")
	}
}
