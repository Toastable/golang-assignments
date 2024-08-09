package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var cityList = [7]string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

func main() {
	const citiesFileName string = "myfavouritecities.txt"

	WriteFileToDisk(cityList[:], citiesFileName)

	fileContents := ReadFileFromDisk(citiesFileName)
	orderedContents := getOrderedFileContents(fileContents)
	fmt.Println(orderedContents)
}

func getOrderedFileContents(contents string) []string {
	orderedContents := strings.Split(contents, "\n")
	sort.Strings(orderedContents)
	return orderedContents
}

func ReadFileFromDisk(filePath string) string {
	contentBytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(contentBytes)
}

func WriteFileToDisk(contents []string, filePath string) {

	if contents == nil {
		contents = cityList[:]
	}

	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for _, city := range contents {
		_, err := file.WriteString(city + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
