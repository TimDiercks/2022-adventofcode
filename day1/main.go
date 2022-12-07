package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	calories := readInput("input.txt")

	maxCaloriesElf, _ := getMaxCaloriesElf(calories)
	max3Sum := getMax3Sum(calories)

	fmt.Printf("Elf number %d carries the most calories with %d cals. \n", maxCaloriesElf, calories[maxCaloriesElf])
	fmt.Printf("The top 3 Elves carry %d calories. \n", max3Sum)

}

func getMax3Sum(array []int) int {
	highestIndicies := []int{0, 1, 2}

	for index, value := range array {
		if arrayContains(index, highestIndicies) {
			continue
		}
		index2, value2 := findMinWithIndex(highestIndicies, array)
		if value > array[value2] {
			highestIndicies[index2] = index
		}
	}

	overAllSum := 0

	for _, value := range highestIndicies {
		overAllSum += array[value]
	}

	return overAllSum
}

func findMinWithIndex(indicies []int, array []int) (int, int) /* (index, value) */ {
	minIndex := 0
	for index, value := range indicies {
		if array[value] < array[indicies[minIndex]] {
			minIndex = index
		}
	}
	return minIndex, indicies[minIndex]
}

func arrayContains(input int, array []int) bool {
	for _, value := range array {
		if input == value {
			return true
		}
	}
	return false
}

func getMaxCaloriesElf(array []int) (int, int) /*(index, value)*/ {
	maxCaloriesElf := 0
	for index, value := range array {
		if value > array[maxCaloriesElf] {
			maxCaloriesElf = index
		}
	}

	return maxCaloriesElf, array[maxCaloriesElf]
}

func readInput(inputFile string) []int { // reads the input file and returns the sum of calories for every elf
	var calories []int

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	currentElf := 0

	// Read Sum of each elf into array
	for fileScanner.Scan() {
		value, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			currentElf++
			continue
		}
		if len(calories) <= currentElf {
			calories = append(calories, 0)
		}
		calories[currentElf] += value
	}

	return calories
}
