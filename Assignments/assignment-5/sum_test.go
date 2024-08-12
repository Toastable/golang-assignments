package main

import (
	"log"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("returns sum of single digits", func(t *testing.T) {
		sut := []int{6, 9, 2}
		want := 17
		got, _ := Sum(sut, 1)

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns error when digit limit is exceeded for single digit", func(t *testing.T) {
		sut := []int{6, 99, 2}
		_, err := Sum(sut, 1)

		if err == nil {
			log.Fatal("Error should have due to digit limit being exceeded")
		}
	})

	t.Run("returns error when length limit is exceeded", func(t *testing.T) {
		sut := []int{6, 9, 2, 5, 7, 8}
		_, err := Sum(sut, 1)

		if err == nil {
			log.Fatal("Error should have due to length limit being exceeded")
		}
	})

	t.Run("returns sum of double digits", func(t *testing.T) {
		sut := []int{66, 99, 22}
		want := 187
		got, err := Sum(sut, 2)

		if err != nil {
			log.Fatal(err)
		}

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns error when digit limit is exceeded for double digit", func(t *testing.T) {
		sut := []int{66, 99, 222}
		_, err := Sum(sut, 2)

		if err == nil {
			log.Fatal("Error should have due to digit limit being exceeded")
		}
	})

	t.Run("returns sum of triple digits", func(t *testing.T) {
		sut := []int{666, 999, 222}
		want := 1887
		got, err := Sum(sut, 3)

		if err != nil {
			log.Fatal(err)
		}

		assertCorrectMessage(t, got, want)
	})

	t.Run("returns error when digit limit is exceeded for triple digit", func(t *testing.T) {
		sut := []int{6666, 999, 222}
		_, err := Sum(sut, 3)

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
