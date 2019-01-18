// https://github.com/gocolly/colly/tree/master/_examples
package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		//colly.AllowedDomains("https://www.zhihu.com/"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"),
		)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	fmt.Println("Launching Scraper !")
	c.Visit("https://www.zhihu.com/people/cuishite/activities")
}

