package index

import (
	"goSearch/pkg/crawler"
	"strings"
)

type Index struct {
	index map[string][]crawler.Document
}

func New() *Index {
	return &Index{
		index: make(map[string][]crawler.Document),
	}
}

func (i *Index) Add(word string, docs []crawler.Document) []crawler.Document {
	for _, d := range docs {
		if strings.Contains(strings.ToLower(d.URL), strings.ToLower(word)) || strings.Contains(strings.ToLower(d.Title), strings.ToLower(word)) {
			i.index[word] = append(i.index[word], d)
		}
	}

	return i.index[word]
}

func (i *Index) Get(word string) []crawler.Document {
	return i.index[strings.ToLower(word)]
}
