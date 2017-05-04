package crawler

import(
"net/http"
"net/url"
"io/ioutil"
"log"
"path/filepath"
"os"
"strings"
"fmt"
"time"
)

/* Generic Helper Methods */

func logger(message string, value interface{}){
	log.Println("|| ", message, value)
}

func delaySeconds(seconds int) {
    time.Sleep(time.Duration(seconds) * time.Second)
}

func chopStringArray(list []string, size_limit int) [][]string{
	var divided [][]string
	chunkSize := (len(list) + size_limit - 1) / size_limit
	for i := 0; i < len(list); i += chunkSize {
		end := i + chunkSize
		if end > len(list) {
			end = len(list)
		}
		divided = append(divided, list[i:end])
	}
	return divided
}

func storageInfoFromURL(url_string string) []string{
	fileName := ""
	urlObject, _ := url.Parse(url_string)
	if (urlObject.Path == "") {
		return []string{urlObject.Host, "index.html"}
	}
	tokens := strings.Split(urlObject.Path, "/")
	fileDirPath := strings.Join(tokens[:len(tokens)-1], "/")
	lastToken := tokens[len(tokens)-1]
	if strings.Contains(lastToken, "."){
		fileName = lastToken
	}else{
		fileName = strings.Join([]string{lastToken, "html"}, ".") 
	}
	return []string{fileDirPath, fileName}
}

/* Crawler Helper Methods */

func writePage(fileName string, content []byte) error{
	logger("[Processing] Writing to the file : ", fileName)
	err := ioutil.WriteFile(fileName, content, 0666)
	if err != nil{
		return err
	}
	return nil
}


func getPageContent(url string) ([]byte, error){
	logger("[Processing] Fetching page content : ", url)
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	return body, nil
}

/* Crawler Object Type */

type Crawler struct{
	MaxConcurrencyLimit int
	CrawlDelay int
	StoragePath  string
	Failures []string
}

/* Crawler Object Methods */

func (config *Crawler) CleanStoragePath(){
	tpath := strings.TrimSpace(config.StoragePath)
	if tpath != config.StoragePath{
		config.StoragePath = tpath
	}
}

func (config *Crawler) CrawlPage(fqdn string, path string, filename string) error{
	config.CleanStoragePath() //cleans configured storage path.
	config_dirpath := filepath.Dir(config.StoragePath)
	dirpath := filepath.Join(config_dirpath, path)
	fullpath := filepath.Join(dirpath, filename)
	err := os.MkdirAll(dirpath, 0777)
	if err != nil{
		return err
	}
	_, err = url.Parse(fqdn)
	if err != nil{
		return err
	}
	pageContent, err := getPageContent(fqdn)
	if err != nil{
		return err
	}
	err = writePage(fullpath, pageContent)
	if err != nil{
		return err
	}
	return nil
}

func (config *Crawler) CrawlPages(url_list []string, done chan<- bool) bool{
	failed_urls := make([]string, len(url_list))
	for _, url := range url_list{
		file_info := storageInfoFromURL(url)
		err := config.CrawlPage(url, file_info[0], file_info[1])
		if err != nil{
			logger("[Error] Failed crawling : ", err)
			config.Failures = append(config.Failures, url)
		}else{
			logger("[Success] Crawled page : ", url)
		}
		delaySeconds(config.CrawlDelay)
		done <- true
	}
	if len(failed_urls) > 0{
		return false
	}
	return true
}



func (config *Crawler) Crawl(url_list []string) bool{
	concurrency_url_lists := [][]string{url_list}
	if config.MaxConcurrencyLimit > 0{
		concurrency_url_lists = chopStringArray(url_list, config.MaxConcurrencyLimit)
	}
	logger("[Processing] Spawning subroutines : ", len(concurrency_url_lists))
	done := make(chan bool)
	for i := 0; i<len(concurrency_url_lists); i++{
		go config.CrawlPages(concurrency_url_lists[i], done)
	}
	for urls_processed := 0; urls_processed<len(url_list); urls_processed++{
		<-done
	}
	close(done)
	logger("[Status] Failed urls : ", config.Failures)
	return true
}