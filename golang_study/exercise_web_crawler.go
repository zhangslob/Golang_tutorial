package main

import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func CrawlResursive(url string, depth int, fetcher Fetcher, quit chan bool, visitedUrls map[string]bool)  {
	if depth <= 0 {
		quit <- true
		return
	}

	didIt, hasIt := visitedUrls[url]
	if didIt && hasIt {
		quit <- true
		return
	} else {
		visitedUrls[url] = true
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		quit <- true
		return
	}
	fmt.Printf("found urls: %s; title: %q\n", url, body)

	childrenQuit := make(chan bool)
	for _, childrenUrl := range urls {
		go CrawlResursive(childrenUrl, depth-1, fetcher, childrenQuit, visitedUrls)
		<- childrenQuit
	}
	quit <- true
}

func Crawl(url string, depth int, fetcher Fetcher)  {
	quit := make(chan bool)
	visitedUrls := map[string]bool{url: false}

	go CrawlResursive(url, depth, fetcher, quit, visitedUrls)

	<-quit
}

func main() {
	Crawl("http://golang.org", 4, fetcher)
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org",
			"http://golang.org/pkg/",
		},
	},
}