package day3

import (
	"fmt"
	"strings"
)

func Solve2(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for i := 0; i < len(lines); i = i + 3 {
		m := map[string]int{}
		t := map[string]int{}
		for _, v := range lines[i] {
			s := string(v)
			if _, ok := m[s]; !ok {
				m[s] = 1
			}
		}
		for _, v := range lines[i+1] {
			s := string(v)
			if v, ok := m[s]; ok {
				if _, ok := t[s]; !ok {
					m[s] = v + 1
					t[s] = 1
				}
			}
		}
		t = map[string]int{}
		for _, v := range lines[i+2] {
			s := string(v)
			if _, ok := t[s]; !ok {
				t[s] = 1
				if v, ok := m[s]; ok {
					m[s] = v + 1
					if m[s] == 3 {
						sum += getP(s)
					}
				}
			}

		}
		m = map[string]int{}
		t = map[string]int{}
	}

	fmt.Println("Answer is", sum)
}
