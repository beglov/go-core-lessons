package storage

// Хранилище отсканированных документов.

import (
	"go-core-lessons/lesson-13/pkg/crawler"
)

// Interface определяет контракт хранилища данных.
type Interface interface {
	Docs([]int) []crawler.Document
	StoreDocs([]crawler.Document) error
	Clear()
}
