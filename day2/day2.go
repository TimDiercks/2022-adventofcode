package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	POINTS_ROCK    = 1
	POINTS_PAPER   = 2
	POINTS_SCISSOR = 3

	POINTS_LOSS = 0
	POINTS_DRAW = 3
	POINTS_WIN  = 6

	ROCK    = "Rock"
	PAPER   = "Paper"
	SCISSOR = "Scissor"
)

func main() {
	opponentsChoice, ownChoice := readInput("input.txt")
	score := computeGameScore(ownChoice, opponentsChoice)
	fmt.Printf("You've got %d points after your game.\n", score)
}

func computeGameScore(ownChoices, opponentsChoices []string) int {
	var score int
	index := 0

	for index < len(ownChoices) {
		ownChoice, err := inputToMove(ownChoices[index])
		if err != nil {
			panic(err)
		}
		opponentsChoice, err := inputToMove(opponentsChoices[index])
		if err != nil {
			panic(err)
		}
		score += computeRoundScore(ownChoice, opponentsChoice)
		index++
	}

	return score
}

func computeRoundScore(ownChoice, opponentsChoice string) int {
	score := computeChoiceScore(ownChoice)
	if ownChoice == ROCK {
		if opponentsChoice == ROCK {
			score += POINTS_DRAW
		}
		if opponentsChoice == PAPER {
			score += POINTS_LOSS
		}
		if opponentsChoice == SCISSOR {
			score += POINTS_WIN
		}
	}
	if ownChoice == PAPER {
		if opponentsChoice == ROCK {
			score += POINTS_WIN
		}
		if opponentsChoice == PAPER {
			score += POINTS_DRAW
		}
		if opponentsChoice == SCISSOR {
			score += POINTS_LOSS
		}
	}
	if ownChoice == SCISSOR {
		if opponentsChoice == ROCK {
			score += POINTS_LOSS
		}
		if opponentsChoice == PAPER {
			score += POINTS_WIN
		}
		if opponentsChoice == SCISSOR {
			score += POINTS_DRAW
		}
	}
	return score
}

func computeChoiceScore(choice string) int {
	if choice == ROCK {
		return 1
	}
	if choice == PAPER {
		return 2
	}
	if choice == SCISSOR {
		return 3
	}
	return 0
}

func readInput(inputFile string) ([]string, []string) {

	var opponentsChoices []string
	var ownChoices []string

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Split(line, " ")
		opponentsChoices = append(opponentsChoices, splitLine[0])
		ownChoices = append(ownChoices, splitLine[1])
	}

	return opponentsChoices, ownChoices
}

func inputToMove(input string) (string, error) {
	if input == "A" || input == "X" {
		return ROCK, nil
	}

	if input == "B" || input == "Y" {
		return PAPER, nil
	}
	if input == "C" || input == "Z" {
		return SCISSOR, nil
	}
	return "nil", fmt.Errorf("Input is invalid")
}
