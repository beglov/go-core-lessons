package engine

import (
	"go-core-lessons/lesson-14/pkg/crawler"
	"go-core-lessons/lesson-14/pkg/crawler/membot"
	"go-core-lessons/lesson-14/pkg/index/hash"
	"go-core-lessons/lesson-14/pkg/storage/btree"
	"os"
	"reflect"
	"testing"
)

var engine *Service

func TestMain(m *testing.M) {
	engine = New(hash.New(), btree.New())
	scanner := membot.New()
	documents, _ := scanner.Scan("http://example.com", 1)
	engine.Add(documents)
	os.Exit(m.Run())
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  []crawler.Document
	}{
		{
			name:  "Тест №1",
			query: "Google",
			want: []crawler.Document{
				{
					ID:    1,
					URL:   "https://google.ru",
					Title: "Google",
				},
			},
		},
		{
			name:  "Тест №2",
			query: "Boogle",
			want:  nil,
		},
		{
			name:  "Тест №3",
			query: "",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := engine.Search(tt.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Clear(t *testing.T) {
	engine.Clear()
	if engine.index.Search("Google") != nil {
		t.Errorf("получили %v, ожидалось %v", engine.index.Search("Google"), nil)
	}
}

func TestService_Add(t *testing.T) {
	engine.Add([]crawler.Document{
		{
			ID:    6,
			URL:   "https://bubble.ru",
			Title: "Bubble",
		},
	})
	got := engine.Search("Bubble")
	want := []crawler.Document{
		{
			ID:    6,
			URL:   "https://bubble.ru",
			Title: "Bubble",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
