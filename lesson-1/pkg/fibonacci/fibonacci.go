package fibonacci

import "log"

// Fibonacci возвращает число Фибоначчи по его порядковому номеру
func Fibonacci(i int) int {
	if i > 20 {
		log.Fatal("число не может быть больше 20")
	}

	if i == 1 || i == 2 {
		return 1
	}

	return Fibonacci(i-1) + Fibonacci(i-2)
}
