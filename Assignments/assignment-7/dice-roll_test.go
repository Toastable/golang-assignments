package main

import (
	"testing"
)

type DiceRollTest struct {
	CombinedDiceRoll int
	ExpectedOutcome  string
}

var DiceRollTests = []DiceRollTest{
	{2, "Snake-Eyes-Craps"},
	{3, "Loss-Craps"},
	{4, "Neutral"},
	{5, "Neutral"},
	{6, "Neutral"},
	{7, "Natural"},
	{8, "Neutral"},
	{9, "Neutral"},
	{10, "Neutral"},
	{11, "Natural"},
	{12, "Loss-Craps"},
}

func TestGenerateDiceOutcome(t *testing.T) {
	for _, sut := range DiceRollTests {
		got := GenerateDiceOutcome(sut.CombinedDiceRoll)

		assertCorrectMessage(t, got, sut.ExpectedOutcome)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
