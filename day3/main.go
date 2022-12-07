package main

import (
	"bufio"
	"fmt"
	"os"
)

type compartment struct {
	items []byte
}

type backpack struct {
	compartments []compartment
}

func main() {
	backpackContents := readInput("input.txt")
	backpacks := createBackpacks(backpackContents)
	sum := computeSum(backpacks)
	fmt.Printf("Solution for Task 1 is: %d\n", sum)

	groupItems := findGroupItems(backpacks)
	sum = 0
	for _, item := range groupItems {
		sum += charToPriority(item)
	}
	fmt.Printf("Solution for Task 2 is: %d\n", sum)
}

func findGroupItems(backpacks []backpack) []byte {
	var groupItems []byte

	for group := 0; group < len(backpacks)/3; group++ {
		firstMemberItems := append(backpacks[group*3].compartments[0].items, backpacks[group*3].compartments[1].items...)
		secondMemberItems := append(backpacks[group*3+1].compartments[0].items, backpacks[group*3+1].compartments[1].items...)
		thirdMemberItems := append(backpacks[group*3+2].compartments[0].items, backpacks[group*3+2].compartments[1].items...)

		groupItems = append(groupItems, getCommonItem(firstMemberItems, secondMemberItems, thirdMemberItems))
	}

	return groupItems
}

func getCommonItem(first, second, third []byte) byte {
	for _, firstItem := range first {
		for _, secondItem := range second {
			if firstItem != secondItem {
				continue
			}
			for _, thirdItem := range third {
				if firstItem == thirdItem {
					return firstItem
				}
			}
		}
	}
	return 0
}

func computeSum(backpacks []backpack) int {
	var summedPriority int

	for _, backpack := range backpacks {
		summedPriority += charToPriority(findSimilarityInBackpack(backpack))
	}

	return summedPriority
}

func findSimilarityInBackpack(backpack backpack) byte {
	for _, first := range backpack.compartments[0].items {
		for _, second := range backpack.compartments[1].items {
			if second == first {
				return first
			}
		}
	}
	return 0
}

func charToPriority(input byte) int {
	if input < 97 {
		return int(input) - 38
	}
	return int(input) - 96
}

func splitIntoCompartments(items string) (compartment, compartment) {
	var first compartment
	var second compartment

	length := len(items)

	for i := 0; i < length; i++ {
		if i < length/2 {
			first.items = append(first.items, items[i])
			continue
		}
		second.items = append(second.items, items[i])
	}

	return first, second
}

func createBackpacks(backpackContents []string) []backpack {
	var backpacks []backpack

	for _, content := range backpackContents {
		var backpack backpack
		firstCompartment, secondCompartment := splitIntoCompartments(content)

		backpack.compartments = append(backpack.compartments, firstCompartment, secondCompartment)
		backpacks = append(backpacks, backpack)
	}

	return backpacks
}

func readInput(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var backpackContents []string

	for fileScanner.Scan() {
		backpackContent := fileScanner.Text()
		backpackContents = append(backpackContents, backpackContent)
	}

	return backpackContents
}
