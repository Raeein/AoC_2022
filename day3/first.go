package day3

import (
	"fmt"
	"strings"
	"unicode"
)

func Solve1(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	m := map[string]int{}
	for _, l := range lines {
		c1 := l[:len(l)/2]
		c2 := l[len(l)/2:]
		for _, c := range c1 {
			s := string(c)
			if v, ok := m[s]; ok {
				m[s] = v + 1
			} else {
				m[s] = 1
			}
		}
		for _, c := range c2 {
			s := string(c)
			if _, ok := m[s]; ok {
				delete(m, s)
				sum += getP(s)
			}
		}
		m = map[string]int{}
	}
	fmt.Println("Answer is", sum)
}

func getP(s string) (p int) {
	r := []rune(s)
	if unicode.IsLower(r[0]) {
		p = int(r[0]-'0') - 48
	} else {
		p = int(r[0]-'0') + 10
	}
	return
}
