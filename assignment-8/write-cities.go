package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	const citiesFileName string = "myfavouritecities.txt"
	var cityList = [7]string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

	WriteFileToDisk(cityList[:], citiesFileName)

	OpenFileFromDisk(citiesFileName)
}

func OpenFileFromDisk(fileName string) {
	contentBytes, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	stringContent := string(contentBytes)
	fmt.Println(stringContent)
}

func WriteFileToDisk(contents []string, fileName string) {
	file, err := os.Create(fileName)

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
