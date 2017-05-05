package web

import(
"os"
"testing"
)

// Tests the crawler with proper configurations and 
// looks for the availability of files downloaded.
func TestCrawlerWithConfigurations(t *testing.T){
	newCrawler := Crawler{
		MaxConcurrencyLimit: 2, 
		StoragePath: "crawler/storage/", 
		CrawlDelay: 1,
	}
	urls := []string{"https://example.com"}
	newCrawler.Start(urls)
	expectedStoragePath := "./crawler/storage/example.com"
	_, err := os.Stat(expectedStoragePath)
	defer os.RemoveAll("./crawler")
	if err != nil{
		t.Error("[Error] Failed to find file on storage path : ", err)
	}
}

// Tests the crawler without configuration which expects 
// to take defaults and looks for the availability of
// files downloaded.
func TestCrawlerWithoutConfigurations(t *testing.T){
	newCrawler := Crawler{}
	urls := []string{"https://example.com"}
	newCrawler.Start(urls)
	expectedStoragePath := "./example.com"
	_, err := os.Stat(expectedStoragePath)
	defer os.RemoveAll("./crawler")
	if err != nil{
		t.Error("[Error] Failed to find file on storage path : ", err)
	}
}