package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputString := readInput("input.txt")
	fmt.Printf("Solution for Task 1: %d\n", processedCharacters(inputString, 4))
	fmt.Printf("Solution for Task 2: %d\n", processedCharacters(inputString, 14))
}

func processedCharacters(input string, uniquelength int) int {
	var count int

	var last4Letters []byte

	for _, char := range input {
		last4Letters = append(last4Letters, byte(char))
		if len(last4Letters) > uniquelength {
			last4Letters = last4Letters[1:]
		}
		count++
		if len(last4Letters) == uniquelength && areElementsUnique(last4Letters) {
			break
		}
	}

	return count
}

func areElementsUnique(in []byte) bool {
	for index, value := range in {
		for index2, value2 := range in {
			if index == index2 {
				continue
			}
			if value == value2 {
				return false
			}
		}
	}
	return true
}

func readInput(inputFile string) string {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		return fileScanner.Text()
	}
	return ""
}
