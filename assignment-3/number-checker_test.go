package main

import "testing"

func TestGenerateFullNameMessage(t *testing.T) {
	t.Run("returns message in range when message is in valid range", func(t *testing.T) {
		got := GenerateNumberRangeMessage("5")
		want := "5 is between 1 and 10\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns message not in range when message is not in valid range", func(t *testing.T) {
		got := GenerateNumberRangeMessage("11")
		want := "11 is not between 1 and 10\n"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
