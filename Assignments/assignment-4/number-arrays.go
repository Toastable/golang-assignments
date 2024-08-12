package main

import (
	"fmt"
	"sort"
)

func main() {
	numberArray := [10]int{2, 1, 4, 10, 5, 9, 7, 8, 6, 3}

	fmt.Println("In Ascending Order:")
	displayArrayAsc(numberArray[:])

	fmt.Println("|-------------------------|")

	fmt.Println("In Descending Order:")
	displayArrayDesc(numberArray[:])

	fmt.Println("|-------------------------|")

	fmt.Println("Even & Odd Sequential:")
	evenOddArrayCount(numberArray[:])
}

func displayArrayAsc(numberArray []int) {
	SortAscending(&numberArray)
	fmt.Println(numberArray)
}

func displayArrayDesc(numberArray []int) {
	SortDescending(&numberArray)
	fmt.Println(numberArray)
}

func evenOddArrayCount(numberArray []int) {
	var evenNumbers, oddNumbers []int
	sortedDescArray := make([]int, len(numberArray))
	copy(sortedDescArray, numberArray)

	sortedAscArray := make([]int, len(numberArray))
	copy(sortedAscArray, numberArray)

	SortAscending(&sortedAscArray)
	SortDescending(&sortedDescArray)

	fmt.Println("Sorted Ascending:")
	evenNumbers, oddNumbers = FindOddAndEvenNumbers(sortedAscArray)
	printEvenAndOdd(evenNumbers, oddNumbers)

	fmt.Println("Sorted Descending:")
	evenNumbers, oddNumbers = FindOddAndEvenNumbers(sortedDescArray)
	printEvenAndOdd(evenNumbers, oddNumbers)
}

func printEvenAndOdd(evenNumbers []int, oddNumbers []int) {
	fmt.Print("Even: ")
	fmt.Println(evenNumbers)
	fmt.Print("Odd: ")
	fmt.Println(oddNumbers)
}

func SortAscending(numberArray *[]int) {
	sort.Ints(*numberArray)
}

func SortDescending(numberArray *[]int) {
	sort.Sort(sort.Reverse(sort.IntSlice(*numberArray)))
}

func FindOddAndEvenNumbers(numberArray []int) ([]int, []int) {
	var evenNumbers, oddNumbers []int

	for _, number := range numberArray {
		if isNumberOdd(number) {
			oddNumbers = append(oddNumbers, number)
		} else {
			evenNumbers = append(evenNumbers, number)
		}
	}

	return evenNumbers, oddNumbers
}

func isNumberOdd(number int) bool {
	return number%2 != 0
}
