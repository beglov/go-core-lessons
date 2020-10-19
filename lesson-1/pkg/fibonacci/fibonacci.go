package fibonacci

// Num возвращает число Фибоначчи по его порядковому номеру
func Num(i int) int {
	if i == 1 || i == 2 {
		return 1
	}

	return Num(i-1) + Num(i-2)
}
