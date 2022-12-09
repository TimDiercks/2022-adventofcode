package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	treeMap := readInput("input.txt")
	_ = treeMap
	fmt.Printf("Solution for task 1: %d\n", countVisibleTrees(treeMap))
	fmt.Printf("Solution for task 2: %d\n", computeLargesViewingDistance(treeMap))
}

func computeLargesViewingDistance(treeMap [][]int) int {
	var largestView int

	for column := 0; column < len(treeMap); column++ {
		for row := 0; row < len(treeMap[column]); row++ {
			currentViewDist := computeViewDistance(treeMap, row, column)
			if currentViewDist > largestView {
				largestView = currentViewDist
			}
		}
	}

	return largestView
}

func computeViewDistance(treeMap [][]int, x, y int) int {
	return computeViewDistanceDirection(treeMap, x, y, 0, 1) *
		computeViewDistanceDirection(treeMap, x, y, 1, 0) *
		computeViewDistanceDirection(treeMap, x, y, 0, -1) *
		computeViewDistanceDirection(treeMap, x, y, -1, 0)
}
func computeViewDistanceDirection(treeMap [][]int, x, y, xDir, yDir int) int {
	var viewDistance int

	treeHeight := treeMap[y][x]
	if yDir == 0 {
		for xCurrent := x + xDir; xCurrent >= 0 && xCurrent < len(treeMap[y]); xCurrent += xDir {
			viewDistance++
			if treeHeight <= treeMap[y][xCurrent] {
				break
			}
		}
	} else {
		for yCurrent := y + yDir; yCurrent >= 0 && yCurrent < len(treeMap); yCurrent += yDir {
			viewDistance++
			if treeHeight <= treeMap[yCurrent][x] {
				break
			}
		}
	}

	return viewDistance
}

func countVisibleTrees(treeMap [][]int) int {
	var count int

	for column := 0; column < len(treeMap); column++ {
		for row := 0; row < len(treeMap[column]); row++ {
			if checkVisibility(treeMap, row, column) {
				count++
			}
		}
	}

	return count
}

func checkVisibility(treeMap [][]int, x, y int) bool {
	if x == 0 || x == len(treeMap[0]) || y == 0 || y == len(treeMap) {
		return true
	}
	if checkVisibilityDirection(treeMap, x, y, 1, 0) {
		return true
	}
	if checkVisibilityDirection(treeMap, x, y, -1, 0) {
		return true
	}
	if checkVisibilityDirection(treeMap, x, y, 0, 1) {
		return true
	}
	if checkVisibilityDirection(treeMap, x, y, 0, -1) {
		return true
	}
	return false
}

func checkVisibilityDirection(treeMap [][]int, x, y, xDir, yDir int) bool {

	treeHeight := treeMap[y][x]
	if yDir == 0 {
		for xCurrent := x + xDir; xCurrent >= 0 && xCurrent < len(treeMap[y]); xCurrent += xDir {
			if treeHeight <= treeMap[y][xCurrent] {
				return false
			}
		}
	} else {
		for yCurrent := y + yDir; yCurrent >= 0 && yCurrent < len(treeMap); yCurrent += yDir {
			if treeHeight <= treeMap[yCurrent][x] {
				return false
			}
		}
	}

	return true
}

func readInput(input string) [][]int {

	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	var treeMap [][]int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var innerMap []int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			innerMap = append(innerMap, num)
		}
		treeMap = append(treeMap, innerMap)
	}

	return treeMap
}
