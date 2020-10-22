package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"go-core-lessons/lesson-2/pkg/spider"
)

func main() {
	var sFlag = flag.String("s", "", "слово для поиска")
	flag.Parse()

	urls := []string{
		"https://go.dev",
		"http://htmlbook.ru",
	}

	fmt.Println("Идет индексирование...")

	totalData := make(map[string]string)
	for _, url := range urls {
		data, err := spider.Scan(url, 2)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
			continue
		}
		for k, v := range data {
			totalData[k] = v
		}
	}

	if *sFlag != "" {
		search(totalData, *sFlag)
		return
	}

	var word string
	for {
		fmt.Print("Введите слово для поиска или exit для выхода: ")
		fmt.Scanln(&word)
		if word == "exit" {
			break
		}
		search(totalData, word)
	}
}

func search(data map[string]string, word string) {
	for k, v := range data {
		if strings.Contains(k, word) || strings.Contains(v, word) {
			fmt.Printf("%s - %s\n", k, v)
		}
	}
}
