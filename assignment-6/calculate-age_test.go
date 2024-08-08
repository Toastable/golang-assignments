package main

import (
	"log"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("returns correct age ", func(t *testing.T) {
		sut := "01/01/1970"
		want := 54
		got, _ := CalculateAge(sut)

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns error when invalid input is provided", func(t *testing.T) {
		sut := "NotADate"
		_, err := CalculateAge(sut)

		if err == nil {
			log.Fatal("Error should have due to digit limit being exceeded")
		}
	})

	t.Run("returns error when invalid separators are provided", func(t *testing.T) {
		sut := "01-01-1970"
		_, err := CalculateAge(sut)

		if err == nil {
			log.Fatal("Error should have due to digit limit being exceeded")
		}
	})

	t.Run("returns error when invalid date format is provided", func(t *testing.T) {
		sut := "30/01/1970"
		_, err := CalculateAge(sut)

		if err == nil {
			log.Fatal("Error should have due to digit limit being exceeded")
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
