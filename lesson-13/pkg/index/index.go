package index

// Обратный индекс отсканированных документов.

import "go-core-lessons/lesson-13/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
	Clear()
}
