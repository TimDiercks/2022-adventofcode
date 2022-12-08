package main

import (
	"bufio"
	"day5/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ship struct {
	stacks []utils.Stack
}

type moveOperation struct {
	amount int
	from   int
	to     int
}

func main() {
	ship, moveOperations := readInput("input.txt")
	ship = moveItems(ship, moveOperations)
	fmt.Printf("Solution for Task 1: %s\n", getTopStackString(ship))
}

func getTopStackString(ship ship) string {
	var result string

	for _, stack := range ship.stacks {
		item, err := stack.Top()
		checkError(err)
		result += string(item)
	}
	return result
}

func moveItems(ship ship, moves []moveOperation) ship {
	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			item, err := ship.stacks[move.from].Pop()
			checkError(err)
			ship.stacks[move.to].Push(item)
		}
	}
	return ship
}

func readInput(inputFile string) (ship, []moveOperation) {

	file, err := os.Open(inputFile)
	checkError(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var ship ship
	for i := 0; i < 9; i++ {
		ship.stacks = append(ship.stacks, utils.Stack{})
	}
	var moveOperations []moveOperation

	readShip(fileScanner, ship)
	moveOperations = readMoveOperations(fileScanner, moveOperations)

	return ship, moveOperations
}

func readShip(fileScanner *bufio.Scanner, ship ship) {
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}
		if line[1] == '1' {
			continue
		}
		for i := 0; i < 9; i++ {
			item := line[i*4+1]
			if item == ' ' {
				continue
			}
			ship.stacks[i].Push(item)
		}
	}
	for i := 0; i < 9; i++ {
		ship.stacks[i].Reverse()
	}
}

func readMoveOperations(fileScanner *bufio.Scanner, moveOperations []moveOperation) []moveOperation {
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitted := strings.Split(line, " ")
		amount, err := strconv.Atoi(splitted[1])
		checkError(err)
		from, err := strconv.Atoi(splitted[3])
		checkError(err)
		to, err := strconv.Atoi(splitted[5])
		checkError(err)
		moveOperations = append(moveOperations,
			moveOperation{
				amount: amount,
				from:   from - 1,
				to:     to - 1,
			}) // converted to arrayID
	}
	return moveOperations
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
