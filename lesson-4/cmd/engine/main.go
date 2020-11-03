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

	idx := index.New()
	idx.Add(documents)

	fmt.Println(documents)
	fmt.Println(idx)

	if *sFlag != "" {
		results := search(documents, idx, *sFlag)
		for _, v := range results {
			fmt.Printf("%s - %s\n", v.URL, v.Title)
		}
		return
	}

	var word string
	for {
		fmt.Print("Введите слово для поиска или exit для выхода: ")
		fmt.Scanln(&word)
		if word == "exit" {
			break
		}
		results := search(documents, idx, word)
		for _, v := range results {
			fmt.Printf("%s - %s\n", v.URL, v.Title)
		}
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

func search(data []spider.Document, idx *index.Service, word string) (results []spider.Document) {
	ids := idx.Search(word)
	for _, id := range ids {
		if v := binarySearch(data, id); v != -1 {
			results = append(results, data[v])
		}
	}
	return results
}

func binarySearch(data []spider.Document, item int) int {
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
