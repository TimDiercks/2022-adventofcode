package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var calories []int
	file, err := os.Open("input.txt")
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

	maxCaloriesElf := 0
	for index, value := range calories {
		if value > calories[maxCaloriesElf] {
			maxCaloriesElf = index
		}
	}

	fmt.Printf("Elf number %d carries the most calories with %d cals. \n", maxCaloriesElf, calories[maxCaloriesElf])
}
