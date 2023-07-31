package index

// Обратный индекс отсканированных документов.

import "go-core-lessons/lesson-19/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
	Clear()
}
