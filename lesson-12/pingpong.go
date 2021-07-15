package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type game struct {
	scoreboard map[string]int
	limit      int
	ch         chan string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := newGame(3)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go play("Вася", game, wg)
	go play("Борис", game, wg)

	game.ch <- "begin"

	wg.Wait()

	game.printScores()
}

func newGame(limit int) *game {
	g := game{
		limit:      limit,
		scoreboard: make(map[string]int),
		ch:         make(chan string),
	}
	return &g
}

func play(name string, game *game, wg *sync.WaitGroup) {
	defer wg.Done()

	for cmd := range game.ch {
		switch cmd {
		case "begin":
			fmt.Println("Новая подача")
			fmt.Printf("%s: %s\n", name, "ping")

			if game.isWin() {
				fmt.Printf("Отличный удар! Подача за %s.\n", name)
				game.scoreboard[name] += 1
				game.ch <- "stop"
			} else {
				game.ch <- "ping"
			}
		case "ping":
			fmt.Printf("%s: %s\n", name, "pong")

			if game.isWin() {
				fmt.Printf("Отличный удар! Подача за %s.\n", name)
				game.scoreboard[name] += 1
				game.ch <- "stop"
			} else {
				game.ch <- "pong"
			}
		case "pong":
			fmt.Printf("%s: %s\n", name, "ping")

			if game.isWin() {
				fmt.Printf("Отличный удар! Подача за %s.\n", name)
				game.scoreboard[name] += 1
				game.ch <- "stop"
			} else {
				game.ch <- "ping"
			}
		case "stop":
			if game.isDone() {
				close(game.ch)
			} else {
				game.ch <- "begin"
			}
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}

func (g *game) isWin() bool {
	if rand.Intn(100) < 20 {
		return true
	}
	return false
}

func (g *game) isDone() bool {
	for _, v := range g.scoreboard {
		if v >= g.limit {
			return true
		}
	}
	return false
}

func (g *game) printScores() {
	fmt.Println("\nИгра закончена!")
	for player, score := range g.scoreboard {
		fmt.Printf("%s: %d\n", player, score)
	}
}
