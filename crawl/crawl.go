package crawl

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"crawler/indexedwords"

	"github.com/PuerkitoBio/goquery"
)

var urls []string = []string{"https://www.off-the-path.com/en/"}
var urlsCount int = len(urls)

var urlsMap = make(map[string]bool)

func CrawlUrls() {

	timeNow := time.Now().Unix()
	scrapedPageCount := 0

	var wg sync.WaitGroup

	for urlsCount > 0 && time.Now().Unix() < timeNow+10 {
		// scrape(urls[0])

		// urls = urls[1:]
		// urlsCount--

		// scrapedPageCount++

		wg.Add(1)

		go func() {

			if len(urls) > 0 {
				scrape(urls[0])

				urls = urls[1:]
				urlsCount--

				scrapedPageCount++
			}

			wg.Done()
		}()

		time.Sleep(200 * time.Millisecond)

		// wg.Wait()
	}

	wg.Wait()

	fmt.Println("Pages scraped:", scrapedPageCount)
}

func scrape(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)

		return
	}
	defer res.Body.Close()

	if !isHTML(res) {
		fmt.Println("Not an HTML document")

		return
	}

	if res.StatusCode >= 300 {
		fmt.Printf("Status code error: %s\n", res.Status)

		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	// Scrape links in current page and add them to the queue
	doc.Find("a").Each(processAnchorTag)

	// Index title of current page and map them to current url
	indexedwords.SetIndexedWords(strings.Split(doc.Find("title").Text(), " "), url)

	// time.Sleep(100 * time.Millisecond)
}

func processAnchorTag(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")

	if exists {
		if !isUrlValid(href) {
			return
		}

		if urlsMap[href] {
			return
		}

		urls = append(urls, href)
		urlsCount++
		urlsMap[href] = true
	}
}

func isHTML(res *http.Response) bool {
	return strings.Contains(res.Header["Content-Type"][0], "text/html")
}

func isUrlValid(url string) bool {
	return strings.HasPrefix(url, "http")
}
