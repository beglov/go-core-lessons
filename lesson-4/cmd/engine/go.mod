module go-core-lessons/lesson-4/cmd/engine

go 1.15

replace go-core-lessons/lesson-4/pkg/spider v0.0.0 => ../../pkg/spider
replace go-core-lessons/lesson-4/pkg/index v0.0.0 => ../../pkg/index

require (
	go-core-lessons/lesson-4/pkg/spider v0.0.0
	go-core-lessons/lesson-4/pkg/index v0.0.0
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321 // indirect
)
