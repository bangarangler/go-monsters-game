package main

import (
	"fmt"

	"github.com/bangarangler/go-monsters-game/interaction"
)

var (
	currentRound = 0
)

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0 // is currentRound divisible by 3 (no float)

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	fmt.Println(userChoice)
	return ""
}

func endGame() {}
