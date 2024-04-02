package storage

import (
	"fmt"
	"goSearch/pkg/crawler"
	"os"
	"strings"
)

func CheckUrls(urls []string) bool {
	dat, err := os.ReadFile("tmp/search-data/urls.txt")
	if err != nil {
		return false
	}

	storedUrls := strings.Fields(string(dat))
	matchedCount := 0

	if len(storedUrls) != len(urls) {
		return false
	}

	for key, url := range urls {
		if url == storedUrls[key] {
			matchedCount++
		}
	}

	return matchedCount == len(urls)
}

func SaveUrls(urls []string) {

	f, err := os.Create("tmp/search-data/urls.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for _, url := range urls {
		f.WriteString(url + "\n")
	}
	fmt.Println("Urls сохранены в tmp/search-data/urls.txt")
}

func SaveDocs(docs []crawler.Document) {

	f, err := os.Create("tmp/search-data/docs.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for _, doc := range docs {
		f.WriteString(doc.URL + "\n")
	}
	fmt.Println("Docs сохранены в tmp/search-data/docs.txt")
}

func GetDocs() []crawler.Document {

	dat, err := os.ReadFile("tmp/search-data/docs.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	storedDocs := strings.Fields(string(dat))
	var docs []crawler.Document
	for _, doc := range storedDocs {
		docs = append(docs, crawler.Document{
			URL: doc,
		})
	}
	fmt.Println("Документы получены из tmp/search-data/docs.txt")
	return docs

}
