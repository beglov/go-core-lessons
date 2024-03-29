package main

import (
	"encoding/json"
	"go-core-lessons/lesson-16/pkg/api"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"go-core-lessons/lesson-16/pkg/crawler"
	"go-core-lessons/lesson-16/pkg/crawler/spider"
	"go-core-lessons/lesson-16/pkg/engine"
	"go-core-lessons/lesson-16/pkg/index/hash"
	"go-core-lessons/lesson-16/pkg/storage/btree"
)

type gosearch struct {
	engine  *engine.Service
	scanner crawler.Interface

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
	gs.engine = engine.New(hash.New(), btree.New())
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	return &gs
}

// init производит инициализацию
func (gs *gosearch) init() {
	if _, err := os.Stat("prev_search_documents.txt"); err == nil {
		err := gs.restore()
		if err != nil {
			log.Println("не удалось восстановить результаты предыдущего сканирования", err)
			gs.scan()
			return
		}
		go gs.scan()
	} else {
		gs.scan()
	}
}

func (gs *gosearch) restore() error {
	bytes, err := ioutil.ReadFile("prev_search_documents.txt")
	if err != nil {
		return err
	}

	var documents []crawler.Document
	err = json.Unmarshal(bytes, &documents)
	if err != nil {
		return err
	}

	err = gs.engine.Add(documents)
	return err
}

// scan производит сканирование сайтов и индексирование данных.
func (gs *gosearch) scan() {
	var documents []crawler.Document

	log.Println("сканирование сайтов")
	var wg sync.WaitGroup // создаем WaitGroup для ожидания завершения сканирования сайтов
	wg.Add(2)

	chDocs, chErr := gs.scanner.BatchScan(gs.sites, 2, 10)

	go func() {
		defer wg.Done()
		for err := range chErr {
			log.Println("ошибка при сканировании сайта:", err)
		}
	}()

	go func() {
		defer wg.Done()
		id := 0
		for doc := range chDocs {
			doc.ID = id
			id++
			documents = append(documents, doc)
		}
	}()

	wg.Wait()
	log.Println("сканирование сайтов завершено")

	gs.engine.Clear()
	err := gs.engine.Add(documents)
	if err != nil {
		log.Println("ошибка при добавлении документов:", err)
	}

	err = gs.dump(documents)
	if err != nil {
		log.Println("не удалось создать дамп результатов сканирования", err)
	}
}

func (gs *gosearch) dump(documents []crawler.Document) error {
	bytes, err := json.Marshal(documents)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("prev_search_documents.txt", bytes, 0644)
	return err
}

func (gs *gosearch) run() {
	srv := api.New(gs.engine)
	log.Fatal(srv.Start(":8000"))
}
