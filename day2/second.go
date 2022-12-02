package day2

import (
	"fmt"
	"strings"
)

func Solve2(input string) {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		s, c := checkTurn(choices[0], choices[1])
		total += s
		total += getTScore(c)
	}
	fmt.Println(total)
}

func checkTurn(t1 string, t2 string) (int, string) {
	moveMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	if t2 == "Y" {
		return DRAW, moveMap[t1]
	}
	if t2 == "Z" {
		if t1 == "A" {
			return WIN, "Y"
		} else if t1 == "B" {
			return WIN, "Z"
		} else {
			return WIN, "X"
		}
	} else {
		if t1 == "A" {
			return LOSE, "Z"
		} else if t1 == "B" {
			return LOSE, "X"
		} else {
			return LOSE, "Y"
		}
	}
}
