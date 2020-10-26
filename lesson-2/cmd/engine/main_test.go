package main

import (
	"reflect"
	"testing"
)

type testScanner struct{}

func (c *testScanner) Scan(url string, depth int) (map[string]string, error) {
	data := make(map[string]string)
	data["https://go.dev/learn"] = "Getting Started - go.dev"
	return data, nil
}

func Test_scanning(t *testing.T) {
	var s testScanner
	got, _ := scan(&s, "https://go.dev", 2)
	want := make(map[string]string)
	want["https://go.dev/learn"] = "Getting Started - go.dev"
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
