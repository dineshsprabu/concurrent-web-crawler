# Concurrent Web Crawler

Highly configurable crawler with powerful concurrency and better status logging.

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

	// Creating a crawler object with configurations.

	c := Crawler{ 
			MaxConcurrencyLimit: 2, 
			StoragePath: "/crawler/storage", 
			CrawlDelay: 10 //in seconds.
		}

	// List of URLS to be crawled as a string array.

	urls := []string{ 
				"https://httpbin.org/ip", 
				"http://example.com", 
				"https://archive.org/details/opensource_movies"
			}

	// Starting the crawler by passing the list of URLs.

	c.Crawl(urls)
}

```