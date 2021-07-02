package main

import (
	"fmt"

	"github.com/bangarangler/go-monsters-game/actions"
	"github.com/bangarangler/go-monsters-game/interaction"
)

var (
	currentRound = 0
	gameRounds   = []interaction.RoundData{}
)

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
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

	var (
		playerAttackDmg  int
		playerHealValue  int
		monsterAttackDmg int
	)

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	// roundData := interaction.RoundData{
	// 	Action:           userChoice,
	// 	PlayerAttackDmg:  playerAttackDmg,
	// 	PlayerHealValue:  playerHealValue,
	// 	MonsterAttackDmg: monsterAttackDmg,
	// 	PlayerHealth:     playerHealth,
	// 	MonsterHealth:    monsterHealth,
	// }
	roundData := interaction.NewRoundData(userChoice, playerAttackDmg, playerHealValue, monsterAttackDmg, playerHealth, monsterHealth)

	// interaction.PrintRoundStatistics(&roundData)
	roundData.PrintStatistics()

	gameRounds = append(gameRounds, *roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
