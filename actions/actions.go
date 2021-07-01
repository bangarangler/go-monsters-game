package actions

import (
	"math/rand"
	"time"
)

var (
	randSource           = rand.NewSource(time.Now().UnixNano())
	randGenerator        = rand.New(randSource)
	currentMonsterHealth = 100
	currentPlayerHealth  = 100
)

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := 5
	maxAttackValue := 10

	if isSpecialAttack {
		minAttackValue = 10
		maxAttackValue = 20
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue
}

func HealPlayer() {
	minHealValue := 10
	maxHealValue := 20

	healValue := generateRandBetween(minHealValue, maxHealValue)

	healthDiff := 100 - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = 100
	}
}

func AttackPlayer() {
	minAttackValue := 9
	maxAttackValue := 13

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= dmgValue
}

func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}
