package interaction

import "fmt"

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
