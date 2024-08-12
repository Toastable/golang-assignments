package main

import "testing"

func TestGenerateFullNameMessage(t *testing.T) {
	t.Run("returns expected full name message", func(t *testing.T) {
		got := GenerateFullNameMessage("Obi", "Wan", "Kenobi")
		want := "You entered Obi Wan Kenobi"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
