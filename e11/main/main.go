package main

import (
	"fmt"
	"github.com/gophercises/e11"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
