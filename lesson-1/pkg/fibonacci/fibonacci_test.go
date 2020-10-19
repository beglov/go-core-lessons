package fibonacci

import (
	"testing"
)

func Test_Num(t *testing.T) {
	got := Num(8)
	want := 21

	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
