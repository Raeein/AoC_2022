package day5

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	width  = 9
	height = 8
)

var stacks = make([][]string, width)

func Solve1(input string) {
	lines := strings.Split(input, "\n")
	makeStacks(lines[:height])
	for i := 0; i < len(lines)-height-2; i++ {
		p := strings.Split(lines[i+height+2], " ")
		loopsS, firstS, secondS := string(p[1]), string(p[3]), string(p[5])
		loops, _ := strconv.Atoi(loopsS)
		first, _ := strconv.Atoi(firstS)
		second, _ := strconv.Atoi(secondS)
		first, second = first-1, second-1
		for loops > 0 {
			stacks[second] = append(stacks[second], stacks[first][len(stacks[first])-1])
			stacks[first] = stacks[first][:len(stacks[first])-1]
			loops--
		}
	}
	for i := 0; i < width; i++ {
		if len(stacks[i]) <= 0 {
			fmt.Println(" ")
		} else {
			fmt.Print(stacks[i][len(stacks[i])-1])
		}
	}
}

func makeStacks(input []string) {
	for _, s := range input {
		for i := 0; i < len(s); i++ {
			if s[i] == '[' {
				stacks[i/4] = append(stacks[i/4], string(s[i+1]))
				i += 3
			}
		}
	}
	for i := 0; i < len(stacks); i++ {
		func(a []string) {
			for i := len(a)/2 - 1; i >= 0; i-- {
				opp := len(a) - 1 - i
				a[i], a[opp] = a[opp], a[i]
			}
		}(stacks[i])
	}
}
