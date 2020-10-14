package main

import (
	"fmt"

	"go-core-lessons/lesson-1/pkg/fibonacci"
)

func main() {
	fmt.Printf("%d число Фибоначчи: %d\n", 1, fibonacci.Fibonacci(1))
	fmt.Printf("%d число Фибоначчи: %d\n", 2, fibonacci.Fibonacci(2))
	fmt.Printf("%d число Фибоначчи: %d\n", 3, fibonacci.Fibonacci(3))
	fmt.Printf("%d число Фибоначчи: %d\n", 4, fibonacci.Fibonacci(4))
	fmt.Printf("%d число Фибоначчи: %d\n", 5, fibonacci.Fibonacci(5))
	fmt.Printf("%d число Фибоначчи: %d\n", 6, fibonacci.Fibonacci(6))
}
