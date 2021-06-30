package interaction

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialAttackAvailable {
			return "SPECIAL ATTACK"
		}
		fmt.Println("Fetching the user input failed. Please try again.")
	}
	// needed when it was for true if you remove true and leave it as above go
	// figures out whats going on
	// return "" // just to make go happy should never get here
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return userInput, nil
}
