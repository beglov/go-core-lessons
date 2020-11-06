package index

import (
	"go-core-lessons/lesson-4/pkg/spider"
	"strings"
)

type Service struct {
	index map[string][]int
}

func New() *Service {
	s := Service{
		index: make(map[string][]int),
	}
	return &s
}

func (s *Service) Add(documents []spider.Document) {
	for _, document := range documents {
		words := strings.Split(document.Title, " ")
		for _, word := range words {
			s.index[word] = append(s.index[word], document.ID)
		}
	}
}

func (s *Service) Search(word string) []int {
	return s.index[word]
}
