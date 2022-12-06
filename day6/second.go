package day6

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"strings"
)

func Solve2(input string) {
	lines := strings.Split(input, "\n")
	set := hashset.New()
	const m = 14
	for _, line := range lines {
		for i := 0; i < len(line)-m; i++ {
			for k := 0; k < m; k++ {
				set.Add(line[i+k])
			}
			if len(set.Values()) == m {
				fmt.Println(i + m)
				set.Clear()
				break
			}
			set.Clear()
		}
	}
}
