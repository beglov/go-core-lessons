package index

import (
	"go-core-lessons/lesson-4/pkg/spider"
	"strings"
)

type Service struct {
	index map[string][]int
}

func New(documents []spider.Document) *Service {
	s := Service{
		index: make(map[string][]int),
	}
	for _, document := range documents {
		words := strings.Split(document.Title, " ")
		for _, word := range words {
			s.index[word] = append(s.index[word], document.ID)
		}
	}
	return &s
}

func (s *Service) Search(word string) []int {
	return s.index[word]
}
