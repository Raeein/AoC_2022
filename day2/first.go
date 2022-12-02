package day2

import (
	"fmt"
	"strings"
)

const (
	WIN  = 6 // Rock
	DRAW = 3 // Paper
	LOSE = 0 // Scissor
	X    = 1 // Rock
	Y    = 2 // Paper
	Z    = 3 // Scissor
)

func Solve1(input string) {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		total += getTScore(choices[1]) + checkWin(choices[0], choices[1])
	}
	fmt.Println(total)
}

func checkWin(t1 string, t2 string) int {
	moveMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	if t2 == moveMap[t1] {
		return DRAW
	}
	if t1 == "A" {
		if t2 == "Y" {
			return WIN
		} else {
			return LOSE
		}
	} else if t1 == "B" {
		if t2 == "X" {
			return LOSE
		} else {
			return WIN
		}
	} else {
		if t2 == "X" {
			return WIN
		} else {
			return LOSE
		}
	}
}

func getTScore(t string) int {
	if t == "X" {
		return X
	} else if t == "Y" {
		return Y
	} else {
		return Z
	}
}
