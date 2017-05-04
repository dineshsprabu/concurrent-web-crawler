# Concurrent Web Crawler

Highly configurable crawler with powerful concurrency and better status logging.

```
go get github.com/dineshsprabu/concurrent-web-crawler


```go

package main

import(
"github.com/dineshsprabu/concurrent-web-crawler"
)

func main(){
	c := Crawler{MaxConcurrencyLimit: 2, StoragePath: "crawled_pages/", CrawlDelay: 10} // CrawlDelay in seconds.
	urls := []string{"https://httpbin.org/ip", "http://example.com", "https://archive.org/details/opensource_movies"}
	c.Crawl(urls)
}

```