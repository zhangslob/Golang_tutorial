package main

import (
	"encoding/json"
	"flag"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"strings"
)

type comment struct {
	Author 		string	`selector:"a.hnuser"`
	URL 		string	`selector:".age a[href]" attr:"href"`
	Comment 	string 	`selector:".comment"`
	Replies		[]*comment
	depth		int
}

func main() {
	var itemID string
	flag.StringVar(&itemID, "id", "", "hackernews post id")
	flag.Parse()
	itemID = "3"

	if itemID == "" {
		log.Println("Hackernews post id required")
		os.Exit(1)
	}

	comments := make([]*comment, 0)

	c := colly.NewCollector()

	c.OnHTML(".comment-tree tr.athing", func(e *colly.HTMLElement) {
		width, err := strconv.Atoi(e.ChildAttr("td.ind img", "width"))
		if err != nil {
			return
		}

		depth := width / 40
		c := &comment{
			Replies: 	make([]*comment, 0),
			depth: 		depth,
		}

		e.Unmarshal(c)
		c.Comment = strings.TrimSpace(c.Comment[:len(c.Comment)-5])
		if depth == 0 {
			comments = append(comments, c)
			return
		}

		parent := comments[len(comments)-1]
		for i :=0; i<depth-1; i++ {
			parent = parent.Replies[len(parent.Replies)-1]
		}
		parent.Replies = append(parent.Replies, c)
	})

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")

	enc.Encode(comments)
}
