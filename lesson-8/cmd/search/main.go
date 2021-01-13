package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"go-core-lessons/lesson-8/pkg/crawler"
	"go-core-lessons/lesson-8/pkg/crawler/spider"
	"go-core-lessons/lesson-8/pkg/engine"
	"go-core-lessons/lesson-8/pkg/index/hash"
	"go-core-lessons/lesson-8/pkg/storage/btree"
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
	id := 0
	var documents []crawler.Document
	for _, url := range gs.sites {
		data, err := gs.scanner.Scan(url, gs.depth)
		if err != nil {
			continue
		}
		for i := range data {
			data[i].ID = id
			id++
		}
		documents = append(documents, data...)
	}

	gs.engine.Clear()
	err := gs.engine.Add(documents)
	if err != nil {
		log.Println("ошибка при добавлении документов:", err)
	}

	err = gs.dump(documents)
	if err != nil {
		log.Println("не удалось сохранить результаты сканирования", err)
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
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите слово для поиска или exit для выхода: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		}
		results := gs.engine.Search(text)
		for _, v := range results {
			fmt.Printf("%s - %s\n", v.URL, v.Title)
		}
	}
}
