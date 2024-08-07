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
	sortAscending(&numberArray)
	fmt.Println(numberArray)
}

func displayArrayDesc(numberArray []int) {
	sortDescending(&numberArray)
	fmt.Println(numberArray)
}

func evenOddArrayCount(numberArray []int) {
	var evenNumbers, oddNumbers []int
	sortedDescArray := make([]int, len(numberArray))
	copy(sortedDescArray, numberArray)

	sortedAscArray := make([]int, len(numberArray))
	copy(sortedAscArray, numberArray)

	sortAscending(&sortedAscArray)
	sortDescending(&sortedDescArray)

	evenNumbers, oddNumbers = findOddAndEvenNumbers(sortedAscArray)
	printEvenAndOdd(evenNumbers, oddNumbers)

	evenNumbers, oddNumbers = findOddAndEvenNumbers(sortedDescArray)
	printEvenAndOdd(evenNumbers, oddNumbers)
}

func printEvenAndOdd(evenNumbers []int, oddNumbers []int) {
	fmt.Println("Sorted Descending:")
	fmt.Print("Even: ")
	fmt.Println(evenNumbers)
	fmt.Print("Odd: ")
	fmt.Println(oddNumbers)
}

func sortAscending(numberArray *[]int) {
	sort.Ints(*numberArray)
}

func sortDescending(numberArray *[]int) {
	sort.Sort(sort.Reverse(sort.IntSlice(*numberArray)))
}

func findOddAndEvenNumbers(numberArray []int) ([]int, []int) {
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
