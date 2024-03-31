package main

import (
	"encoding/json"
	"fmt"
	"goSearch/pkg/crawler"
	"goSearch/pkg/crawler/spider"
	"goSearch/pkg/index"
	"net/http"
)

// var urls = []string{"https://go.dev", "https://vuejs.org"}
var urls = []string{"https://go.dev"}
var documents = make([]crawler.Document, 0)
var idx *index.Index = index.New()

const port = ":8080"

func main() {
	fmt.Printf("Начинаю сканирование документов %+v", urls)

	err := scan(urls)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Сканирование завершено")

	http.HandleFunc("/", handleFind)
	fmt.Printf("Сервер запущен на порту%v\n", port)
	http.ListenAndServe(port, nil)

}

func scan(urls []string) error {
	s := spider.New()
	id := 0 //счетчик документов
	for _, u := range urls {

		data, err := s.Scan(u, 2)
		if err != nil {
			return err
		}

		for _, d := range data {
			id++
			d.ID = id
			documents = append(documents, d)
		}
	}
	return nil
}

func handleFind(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	search := r.URL.Query().Get("q")
	if search == "" {
		http.Error(w, "Missing query parameter 'q'", http.StatusBadRequest)
		return
	}
	result := findByWord(search)

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func findByWord(word string) []crawler.Document {

	fmt.Printf("Поиск по слову %v\n", word)

	result := idx.Get(word)

	if len(result) > 0 {
		fmt.Println("Слово есть в индексе")
	}
	if len(result) == 0 {
		fmt.Println("Не найдено слово в индексах, добавляю его в индекс")
		result = idx.Add(word, documents)
	}

	return result
}
