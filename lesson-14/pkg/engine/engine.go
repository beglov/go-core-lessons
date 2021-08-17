package engine

import (
	"go-core-lessons/lesson-14/pkg/crawler"
	"go-core-lessons/lesson-14/pkg/index"
	"go-core-lessons/lesson-14/pkg/storage"
)

// Engine - поисковый движок.
// Его задача - обслуживание поисковых запросов.
// функциональность:
// - обработка поискового запроса;
// - поиск документов в индексе;
// - запрос документов из хранилища;
// - добавление документов в хранилище и индекс;
// - очистка хранилища и индекса;
// - возврат посиковой выдачи.

// Service - поисковый движок.
type Service struct {
	index   index.Interface
	storage storage.Interface
}

// New - конструктор.
func New(index index.Interface, storage storage.Interface) *Service {
	s := Service{
		index:   index,
		storage: storage,
	}
	return &s
}

// Search ищет документы, соответствующие поисковому запросу.
func (s *Service) Search(query string) []crawler.Document {
	if query == "" {
		return nil
	}
	ids := s.index.Search(query)
	docs := s.storage.Docs(ids)
	return docs
}

// Clear очищает индекс и хранилище
func (s *Service) Clear() {
	s.index.Clear()
	s.storage.Clear()
}

// Add добавляет документы в хранилице и индексирует их
func (s *Service) Add(docs []crawler.Document) error {
	s.index.Add(docs)
	err := s.storage.StoreDocs(docs)
	return err
}
