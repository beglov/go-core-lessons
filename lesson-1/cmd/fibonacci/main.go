package main

import (
	"flag"
	"fmt"
	"log"

	"go-core-lessons/lesson-1/pkg/fibonacci"
)

func main() {
	var nFlag = flag.Int("n", 1, "порядковый номер числа Фибоначчи")
	flag.Parse()

	if *nFlag > 20 {
		log.Fatal("число не может быть больше 20")
	}

	fmt.Printf("%d число Фибоначчи: %d\n", *nFlag, fibonacci.Num(*nFlag))
}
