package main

import (
	"fmt"
	"math/rand"
	"os"
)

type RollResult struct {
	rolls   [2]int
	outcome string
}

func main() {
	for i := 0; i < 50; i++ {
		diceRoll := RollPairOfDice()

		printDiceRolls(diceRoll)
	}
}

func printDiceRolls(diceRoll RollResult) {
	for i, roll := range diceRoll.rolls {
		fmt.Fprintln(os.Stdout, "Die", i+1, "roll was", roll)
	}

	fmt.Fprintln(os.Stdout, "Outcome was", diceRoll.outcome)
}

func RollPairOfDice() RollResult {
	firstDie := rollDie()
	secondDie := rollDie()

	var rollsArray = [2]int{firstDie, secondDie}
	return RollResult{
		rolls:   rollsArray,
		outcome: GenerateDiceOutcome(firstDie + secondDie),
	}
}

func rollDie() int {
	return rand.Intn(6-1+1) + 1
}

func GenerateDiceOutcome(result int) string {
	switch result {
	case 7, 11:
		return "Natural"
	case 2:
		return "Snake-Eyes-Craps"
	case 3, 12:
		return "Loss-Craps"
	default:
		return "Neutral"
	}
}
