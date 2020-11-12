package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/hash"
	"gosearch/pkg/storage/btree"
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
		gs.restore()
	}
	go gs.scan()
}

func (gs *gosearch) restore() {
	log.Println("Восстановление результатов предыдущего сканирования...")

	bytes, err := ioutil.ReadFile("prev_search_documents.txt")
	if err != nil {
		log.Println("ошибка при чтении предыдущих результатов поиска:", err)
		return
	}

	var documents []crawler.Document
	err = json.Unmarshal(bytes, &documents)
	if err != nil {
		log.Println("ошибка при десериализации предыдущих результатов поиска:", err)
		return
	}

	err = gs.engine.Add(documents)
	if err != nil {
		log.Println("ошибка при добавлении документов:", err)
	}

	log.Println("Восстановление завершено")
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

	gs.dump(documents)
}

func (gs *gosearch) dump(documents []crawler.Document) {
	bytes, err := json.Marshal(documents)
	if err != nil {
		log.Println("ошибка при сериализации результатов поиска:", err)
		return
	}
	err = ioutil.WriteFile("prev_search_documents.txt", bytes, 0644)
	if err != nil {
		log.Println("ошибка при сериализации результатов поиска:", err)
		return
	}
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
