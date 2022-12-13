package main

import (
	"github.com/Raeein/AoC_2022/day7"
	"github.com/Raeein/AoC_2022/files"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arg := os.Args[1]
	lines, err := files.ReadFromFileString(arg)
	check(err)
	day7.Solve1(lines)
}
