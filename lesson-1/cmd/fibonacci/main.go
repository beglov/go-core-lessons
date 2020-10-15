package main

import (
	"flag"
	"fmt"

	"go-core-lessons/lesson-1/pkg/fibonacci"
)

func main() {
	var nFlag = flag.Int("n", 1, "порядковый номер числа Фибоначчи")
	flag.Parse()
	fmt.Printf("%d число Фибоначчи: %d\n", *nFlag, fibonacci.Fibonacci(*nFlag))
}
