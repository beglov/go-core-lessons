package engine

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go-core-lessons/lesson-10/gosearch/pkg/crawler"
	"go-core-lessons/lesson-10/gosearch/pkg/index"
	"go-core-lessons/lesson-10/gosearch/pkg/storage"
)

// Engine - поисковый движок.
// Его задача - обслуживание поисковых запросов.
// функциональность:
// - обработка поискового запроса;
// - поиск документов в индексе;
// - запрос документов из хранилища;
// - возврат посиковой выдачи.

// Service - поисковый движок.
type Service struct {
	index   index.Interface
	storage storage.Interface
}

var queryLen = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "query_len",
	Help:    "Длина поискового запроса, байт.",
	Buckets: prometheus.LinearBuckets(5, 5, 40),
})

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
	queryLen.Observe(float64(len(query)))
	return docs
}
