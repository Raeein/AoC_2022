package day4

import (
	"fmt"
	"strings"
)

func Solve2(input string) {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		nums := stringToInt(line)
		if (nums[0] <= nums[2] && nums[1] >= nums[3]) ||
			(nums[0] >= nums[2] && nums[1] <= nums[3]) ||
			(nums[2] >= nums[0] && nums[2] <= nums[1]) ||
			(nums[3] >= nums[0] && nums[3] <= nums[1]) {
			count++
		}
	}
	fmt.Println("Answer is", count)
}
