package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", e)
	})
	c.Visit("https://definitely-not-a.website/")
}
