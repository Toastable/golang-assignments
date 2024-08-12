package main

import (
	"reflect"
	"testing"
)

func TestSortAscending(t *testing.T) {
	t.Run("returns array in ascending order array when array is already sorted", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := make([]int, len(want))
		copy(got, want[:])
		SortAscending(&got)
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns array in ascending order array when array is not sorted", func(t *testing.T) {
		sut := []int{1, 3, 2, 7, 10, 6, 4, 9, 8, 5}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := make([]int, len(sut))
		copy(got, sut[:])
		SortAscending(&got)
		assertCorrectMessage(t, got, want)
	})
}

func TestSortDescending(t *testing.T) {
	t.Run("returns array in descending order array when array is already sorted", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := make([]int, len(want))
		copy(got, want[:])
		SortAscending(&got)
		assertCorrectMessage(t, got, want)
	})

	t.Run("returns array in descending order array when array is not sorted", func(t *testing.T) {
		sut := []int{1, 3, 2, 7, 10, 6, 4, 9, 8, 5}
		want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		got := make([]int, len(sut))
		copy(got, sut[:])
		SortDescending(&got)
		assertCorrectMessage(t, got, want)
	})
}

func TestFindOddAndEvenNumbers(t *testing.T) {
	t.Run("returns sequentially ordered arrays when array is in ascending order", func(t *testing.T) {
		wantEvenNumbers := []int{2, 4, 6, 8, 10}
		wantOddNumbers := []int{1, 3, 5, 7, 9}
		sut := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := make([]int, len(sut))
		copy(got, sut[:])
		even, odd := FindOddAndEvenNumbers(got)
		assertCorrectMessage(t, odd, wantOddNumbers)
		assertCorrectMessage(t, even, wantEvenNumbers)
	})

	t.Run("returns sequentially ordered arrays when array is in descending order", func(t *testing.T) {
		wantEvenNumbers := []int{10, 8, 6, 4, 2}
		wantOddNumbers := []int{9, 7, 5, 3, 1}
		sut := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		got := make([]int, len(sut))
		copy(got, sut[:])
		even, odd := FindOddAndEvenNumbers(got)
		assertCorrectMessage(t, odd, wantOddNumbers)
		assertCorrectMessage(t, even, wantEvenNumbers)
	})
}

func assertCorrectMessage(t testing.TB, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
