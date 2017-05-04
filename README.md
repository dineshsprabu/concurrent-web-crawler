# Concurrent Web Crawler

Highly configurable crawler with powerful concurrency and better status logging.

[![GoDoc](https://godoc.org/github.com/dineshsprabu/concurrent-web-crawler?status.svg)](https://godoc.org/github.com/dineshsprabu/concurrent-web-crawler)

## Installation

```
go get github.com/dineshsprabu/concurrent-web-crawler

```

## Usage

```go

package main

import(
"github.com/dineshsprabu/concurrent-web-crawler"
)

func main(){
	// Creating a web crawler object with configurations.
	myCrawler := web.Crawler{ 
			MaxConcurrencyLimit: 2, 
			StoragePath: "crawler/storage", 
			CrawlDelay: 10,
		}

	// List of URLS to be crawled as a string array.
	urls := []string{ 
				"https://httpbin.org/ip", 
				"http://example.com", 
				"https://archive.org/details/opensource_movies",
			}

	// Starting the crawler by passing the list of URLs.
	myCrawler.Start(urls)
}

```

## Log

```

> go run crawler_test.go 
2017/05/04 20:29:59 ||  [Processing] Spawning subroutines :  2
2017/05/04 20:29:59 ||  [Processing] Fetching page content :  https://archive.org/details/opensource_movies
2017/05/04 20:29:59 ||  [Processing] Fetching page content :  https://httpbin.org/ip
2017/05/04 20:30:01 ||  [Processing] Writing to the file :  crawler/ip.html
2017/05/04 20:30:01 ||  [Success] Crawled page :  https://httpbin.org/ip
2017/05/04 20:30:03 ||  [Processing] Writing to the file :  crawler/details/opensource_movies.html
2017/05/04 20:30:03 ||  [Success] Crawled page :  https://archive.org/details/opensource_movies
2017/05/04 20:30:11 ||  [Processing] Fetching page content :  http://example.com
2017/05/04 20:30:12 ||  [Processing] Writing to the file :  crawler/example.com/index.html
2017/05/04 20:30:12 ||  [Success] Crawled page :  http://example.com
2017/05/04 20:30:22 ||  [Status] Failed urls :  []

```