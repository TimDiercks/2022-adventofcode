package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start int
	end   int
}

type pair struct {
	first  section
	second section
}

func main() {
	pairs := readInput("input.txt")
	overlappingSections := countFullOverlappings(pairs)
	fmt.Printf("Solution for Task 1 is: %d\n", overlappingSections)

	overlappingSections = countOverlappings(pairs)
	fmt.Printf("Solution for Task 2 is: %d\n", overlappingSections)
}

func countFullOverlappings(pairs []pair) int {
	var overlappings int

	for _, pair := range pairs {
		if checkFullOverlappingForPair(pair) {
			overlappings++
		}
	}

	return overlappings
}

func checkFullOverlappingForPair(pair pair) bool {
	return checkFullOverlapping(pair.first, pair.second) || checkFullOverlapping(pair.second, pair.first)
}

func checkFullOverlapping(first section, second section) bool {
	if first.start <= second.start && first.end >= second.end {
		return true
	}
	return false
}

func countOverlappings(pairs []pair) int {
	var overlappings int

	for _, pair := range pairs {
		if checkOverlappingForPair(pair) {
			overlappings++
		}
	}

	return overlappings
}

func checkOverlappingForPair(pair pair) bool {
	return checkOverlapping(pair.first, pair.second) || checkOverlapping(pair.second, pair.first)
}
func checkOverlapping(first section, second section) bool {
	if (first.start >= second.start && first.start <= second.end) || (first.end >= second.start && first.end <= second.end) {
		return true
	}
	return false
}

func stringToInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func getPairFromLine(splitForElf []string) pair {
	var pair pair
	splittedIdsFirst := strings.Split(splitForElf[0], "-")
	pair.first.start = stringToInt(splittedIdsFirst[0])
	pair.first.end = stringToInt(splittedIdsFirst[1])
	splittedIdsSecond := strings.Split(splitForElf[1], "-")
	pair.second.start = stringToInt(splittedIdsSecond[0])
	pair.second.end = stringToInt(splittedIdsSecond[1])
	return pair
}

func readInput(inputFile string) []pair {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var pairs []pair

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitForElf := strings.Split(line, ",")

		pair := getPairFromLine(splitForElf)

		pairs = append(pairs, pair)
	}

	return pairs
}
