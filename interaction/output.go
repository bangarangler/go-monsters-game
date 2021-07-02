package interaction

import (
	"fmt"
	"os"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func NewRoundData(a string, pad, phv, mad, ph, mh int) *RoundData {
	roundData := RoundData{
		Action:           a,
		PlayerAttackDmg:  pad,
		PlayerHealValue:  ph,
		MonsterAttackDmg: mad,
		PlayerHealth:     ph,
		MonsterHealth:    mh,
	}
	return &roundData
}

func (r *RoundData) PrintStatistics() {
	if r.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", r.PlayerAttackDmg)
	} else if r.Action == "SPECIAL ATTACK" {
		fmt.Printf("Player performed a string attack against monster for %v damage.\n", r.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", r.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v.\n", r.MonsterAttackDmg)
	fmt.Printf("Player Health: %v.\n", r.PlayerHealth)
	fmt.Printf("Monster Health: %v.\n", r.MonsterHealth)
}

func PrintGreeting() {
	fmt.Println("MONSTER GAME")
	fmt.Println("Staring a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackAvailable bool) {
	fmt.Println("Please chosse your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackAvailable {
		fmt.Println("(3) Special Attack")
	}
}

// turned into a method above
// func PrintRoundStatistics(roundData *RoundData) {
// 	if roundData.Action == "ATTACK" {
// 		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
// 	} else if roundData.Action == "SPECIAL ATTACK" {
// 		fmt.Printf("Player performed a string attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
// 	} else {
// 		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
// 	}
//
// 	fmt.Printf("Monster attacked player for %v.\n", roundData.MonsterAttackDmg)
// 	fmt.Printf("Player Health: %v.\n", roundData.PlayerHealth)
// 	fmt.Printf("Monster Health: %v.\n", roundData.MonsterHealth)
// }

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("gamelog.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}
	for idx, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(idx + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)
		if err != nil {
			fmt.Println("Writing into log file failed. Exiting.")
			continue
		}
	}
	fmt.Println("Wrote data to log!")
}
