package main

import "testing"

func TestHelloThere(t *testing.T) {
	t.Run("returns expected string", func(t *testing.T) {
		got := HelloThere()
		want := "Grevious: General Kenobi, Obi Wan: Hello There"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
