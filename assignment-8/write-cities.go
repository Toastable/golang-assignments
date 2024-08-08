package main

import (
	"log"
	"os"
)

func main() {
	var cityList = [7]string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

	WriteFileToDisk(cityList[:])
}

func WriteFileToDisk(contents []string) {
	file, err := os.Create("myfavouritecities.txt")

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
