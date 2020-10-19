module go-core-lessons/lesson-2/cmd/engine

go 1.15

replace go-core-lessons/lesson-2/pkg/spider v0.0.0 => ../../pkg/spider

require (
	go-core-lessons/lesson-2/pkg/spider v0.0.0
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
)
