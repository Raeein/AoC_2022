package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1(input string) {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		nums := stringToInt(line)
		if (nums[0] <= nums[2] && nums[1] >= nums[3]) ||
			nums[0] >= nums[2] && nums[1] <= nums[3] {
			count++
		}
	}
	fmt.Println("Answer is", count)
}

func stringToInt(s string) (r []int) {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ",", " ")
	nums := strings.Split(s, " ")

	for _, v := range nums {
		res, _ := strconv.Atoi(v)
		r = append(r, res)
	}
	return
}
