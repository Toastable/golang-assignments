package main

import (
	"log"
	"testing"
)

func TestGenerateFullNameMessage(t *testing.T) {
	t.Run("returns message in range when message is in valid range", func(t *testing.T) {
		got, _ := GenerateNumberRangeMessage("5")
		want := "5 is between 1 and 10\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns message not in range when message is not in valid range", func(t *testing.T) {
		got, _ := GenerateNumberRangeMessage("11")
		want := "11 is not between 1 and 10\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns error when invalid input is provided", func(t *testing.T) {
		_, err := GenerateNumberRangeMessage("NotANumber")

		if err == nil {
			log.Fatal("Error should have occurred during number conversion")
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
