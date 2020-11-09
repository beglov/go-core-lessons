package main

import (
	"flag"
	"fmt"
	"log"

	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index"
	"gosearch/pkg/index/hash"
	"gosearch/pkg/storage"
	"gosearch/pkg/storage/btree"
)

type gosearch struct {
	engine  *engine.Service
	scanner crawler.Interface
	index   index.Interface
	storage storage.Interface

	sites []string
	depth int
}

func main() {
	server := new()
	server.init()
	server.run()
}

// new создаёт объект и службы сервера и возвращает указатель на него.
func new() *gosearch {
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.index = hash.New()
	gs.storage = btree.New()
	gs.engine = engine.New(gs.index, gs.storage)
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	return &gs
}

// init производит сканирование сайтов и индексирование данных.
func (gs *gosearch) init() {
	log.Println("Сканирование сайтов.")
	id := 0
	for _, url := range gs.sites {
		log.Println("Сайт:", url)
		data, err := gs.scanner.Scan(url, gs.depth)
		if err != nil {
			continue
		}
		for i := range data {
			data[i].ID = id
			id++
		}
		log.Println("Индексирование документов.")
		gs.index.Add(data)
		log.Println("Сохранение документов.")
		err = gs.storage.StoreDocs(data)
		if err != nil {
			log.Println("ошибка при добавлении сохранении документов в БД:", err)
			continue
		}
	}
	fmt.Println(gs.storage)
	fmt.Println(gs.index)
}

func (gs *gosearch) run() {
	var sFlag = flag.String("s", "", "слово для поиска")
	flag.Parse()

	if *sFlag != "" {
		results := gs.engine.Search(*sFlag)
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
		results := gs.engine.Search(word)
		for _, v := range results {
			fmt.Printf("%s - %s\n", v.URL, v.Title)
		}
	}
}
