package day6

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"strings"
)

func Solve1(input string) {
	lines := strings.Split(input, "\n")
	set := hashset.New()
	for _, line := range lines {
		for i := 0; i < len(line)-4; i++ {
			set.Add(line[i], line[i+1], line[i+2], line[i+3])
			if len(set.Values()) == 4 {
				fmt.Println(i + 4)
				set.Clear()
				break
			}
			set.Clear()
		}
	}
}
