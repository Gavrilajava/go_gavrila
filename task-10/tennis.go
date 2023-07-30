package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const limit = 5

type Game struct {
	Score map[string]int
	Ch    chan string
	Wg    *sync.WaitGroup
}

type Player struct {
	Name string
	*Game
}

func main() {
	game := newGame()

	game.Wg.Add(2)

	p1 := Player{"Forrest", game}
	p2 := Player{"Bubba", game}

	go p1.play()
	go p2.play()

	game.Ch <- "begin"

	game.Wg.Wait()

	fmt.Println("Game over.")
	for player, score := range game.Score {
		fmt.Printf("%s: %d\n", player, score)
	}
}

func newGame() *Game {
	g := Game{
		Score: make(map[string]int),
		Ch:    make(chan string),
		Wg:    &sync.WaitGroup{},
	}
	return &g
}

func (p *Player) strike(word string) {
	fmt.Println(p.Name, word)
	if rand.Intn(100) > 80 {
		fmt.Printf("Score!!! %s has the ball.\n", p.Name)
		p.Score[p.Name] += 1
		p.Ch <- "stop"
	} else {
		p.Ch <- word
	}
}

func (p *Player) play() {
	defer p.Wg.Done()

	for cmd := range p.Ch {
		switch cmd {
		case "begin":
			p.strike("ping")
		case "ping":
			p.strike("ping")
		case "pong":
			p.strike("pong")
		case "stop":
			if p.Game.isOver() {
				close(p.Ch)
			} else {
				p.Ch <- "begin"
			}
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}

func (g Game) isOver() bool {
	for _, v := range g.Score {
		if v >= limit {
			return true
		}
	}
	return false
}
