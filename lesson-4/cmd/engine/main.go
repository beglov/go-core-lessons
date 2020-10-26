package main

import (
	"flag"
	"fmt"
	"log"

	"go-core-lessons/lesson-4/pkg/index"
	"go-core-lessons/lesson-4/pkg/spider"
)

// Scanner - интерфейс поискового робота.
type Scanner interface {
	Scan(url string, depth int) ([]spider.Document, error)
}

func main() {
	var sFlag = flag.String("s", "", "слово для поиска")
	flag.Parse()

	urls := []string{
		"https://go.dev",
		"http://htmlbook.ru",
	}

	fmt.Println("Идет индексирование...")

	s := spider.New()
	documents := scan(s, urls, 2)

	idx := index.New(documents)

	fmt.Println(documents)
	fmt.Println(idx)

	if *sFlag != "" {
		search(documents, idx, *sFlag)
		return
	}

	var word string
	for {
		fmt.Print("Введите слово для поиска или exit для выхода: ")
		fmt.Scanln(&word)
		if word == "exit" {
			break
		}
		search(documents, idx, word)
	}
}

func scan(s Scanner, urls []string, depth int) []spider.Document {
	var documents []spider.Document
	var id int
	for _, url := range urls {
		s := spider.New()
		data, err := s.Scan(url, depth)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
			continue
		}
		for i, _ := range data {
			data[i].ID = id
			id = id + 1
		}
		documents = append(documents, data...)
	}
	return documents
}

func search(data []spider.Document, idx *index.Service, word string) {
	ids := idx.Search(word)
	for _, id := range ids {
		if v := binary(data, id); v != -1 {
			fmt.Printf("%s - %s\n", data[v].URL, data[v].Title)
		}
	}
}

func binary(data []spider.Document, item int) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].ID == item {
			return mid
		}
		if data[mid].ID < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
