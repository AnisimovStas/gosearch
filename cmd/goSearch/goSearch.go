package main

import (
	"flag"
	"fmt"
	"goSearch/pkg/crawler"
	"goSearch/pkg/crawler/spider"
	"strings"
)

var urls = []string{"https://go.dev", "https://vuejs.org"}

func main() {
	sf := flag.String("s", "", "word to search")
	flag.Parse()

	if *sf == "" {
		fmt.Println("Для использования программы необходимо указать слово для поиска через флаг '-s' \nНапример: goSearch -s котики")
		return
	}
	fmt.Printf("Ищу ссылки со словом: %s\n", *sf)

	documents, err := scan(urls)
	if err != nil {
		return
	}

	printResult(*sf, documents)
}

func scan(urls []string) ([]crawler.Document, error) {
	results := make([]crawler.Document, 0)
	s := spider.New()

	for _, u := range urls {

		data, err := s.Scan(u, 2)
		if err != nil {
			return nil, err
		}
		results = append(results, data...)
	}
	return results, nil
}

func printResult(w string, r []crawler.Document) {
	counter := 0
	for _, d := range r {
		if strings.Contains(strings.ToLower(d.URL), strings.ToLower(w)) || strings.Contains(strings.ToLower(d.Title), strings.ToLower(w)) {
			fmt.Printf("%+v\n", d)
			counter++
		}
	}
	if counter == 0 {
		fmt.Printf("Результатов для слова %v не найдено\n", w)
	}

}
