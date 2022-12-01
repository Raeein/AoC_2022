package main

import (
	"github.com/Raeein/AoC_2022/day1"
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
	day1.Solve2(lines)
}
