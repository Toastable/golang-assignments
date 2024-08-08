package main

import (
	"reflect"
	"testing"
)

// func TestSortAscending(t *testing.T) {
// 	t.Run("returns array in ascending order array when array is already sorted", func(t *testing.T) {
// 		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 		got := make([]int, len(want))
// 		copy(got, want[:])
// 		SortAscending(&got)
// 		assertCorrectMessage(t, got, want)
// 	})
// }

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
