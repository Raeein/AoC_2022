package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve2(input string) {
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
			if loops > 1 {
				stacks[second] = append(
					stacks[second],
					stacks[first][len(stacks[first])-loops:]...,
				)
				stacks[first] = stacks[first][:len(stacks[first])-loops]
				break
			} else {
				stacks[second] = append(
					stacks[second],
					stacks[first][len(stacks[first])-1],
				)
				stacks[first] = stacks[first][:len(stacks[first])-1]
			}
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
	fmt.Printf("\n")
}
