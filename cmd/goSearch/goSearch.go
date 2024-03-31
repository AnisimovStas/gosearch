package main

import (
	"encoding/json"
	"fmt"
	"goSearch/pkg/crawler"
	"goSearch/pkg/crawler/spider"
	"net/http"
	"strings"
)

var urls = []string{"https://go.dev", "https://vuejs.org"}
var documents = make([]crawler.Document, 0)

const port = ":8080"

func main() {
	fmt.Printf("Начинаю сканирование документов %+v", urls)
	err := scan(urls)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Сканирование завершено")

	http.HandleFunc("/", findByWord)
	fmt.Printf("Сервер запущен на порту%v\n", port)
	http.ListenAndServe(port, nil)

}

func scan(urls []string) error {
	//results := make([]crawler.Document, 0)
	s := spider.New()

	for _, u := range urls {

		data, err := s.Scan(u, 2)
		if err != nil {
			return err
		}
		documents = append(documents, data...)
	}
	return nil
}

func findByWord(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Path[1:]

	result := make([]crawler.Document, 0)

	for _, d := range documents {
		if strings.Contains(strings.ToLower(d.URL), strings.ToLower(search)) || strings.Contains(strings.ToLower(d.Title), strings.ToLower(search)) {
			result = append(result, d)
		}
	}
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
