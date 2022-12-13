package day7

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"strconv"
	"strings"
)

func Solve1(input string) {
	lines := strings.Split(input, "\n")
	_ = hashset.New()

	pwd := make([]string, 0)
	dirs_size := make(map[string]int)

	for _, line := range lines {
		log := strings.Split(line, " ")
		if log[1] == "cd" && log[2] != ".." {
			pwd = append(pwd, log[2])
		} else if log[1] == "cd" && log[2] == ".." {
			pwd = pwd[:len(pwd)-1]
		} else {
			size, err := strconv.Atoi(log[0])
			if err == nil {
				path := ""
				for _, dir := range pwd {
					path += dir + "/"
					dirs_size[path] += size
				}
			}
		}
	}

	disk := 70000000
	need := 30000000
	min_size := dirs_size["//"]
	check_size := need - (disk - dirs_size["//"])
	for dir, size := range dirs_size {
		if dir != "/" && size >= check_size {
			if size < min_size {
				min_size = size
			}
		}
	}
	fmt.Println(min_size)
}
