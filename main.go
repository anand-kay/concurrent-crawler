package main

import (
	"fmt"

	"crawler/crawl"
	"crawler/search"
)

func main() {
	crawl.CrawlUrls()

	fmt.Println(search.SearchWords("for"))
}
