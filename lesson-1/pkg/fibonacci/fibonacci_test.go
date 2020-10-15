package fibonacci

import (
	"testing"
)

func Test_Fibonacci(t *testing.T) {
	got := Fibonacci(8)
	want := 21

	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
