package day1

import (
	"fmt"
	"github.com/Raeein/AoC_2022/sliceHelper"
	"strings"
)

func Solve1(input string) {
	snackGroups := strings.Split(input, "\n\n")
	max := 0
	for _, group := range snackGroups {

		groupCalories := strings.Split(group, "\n")
		groupCaloriesInt := sliceHelper.MakeIntSlice(groupCalories)
		newSum := sliceHelper.SumIntSlice(groupCaloriesInt...)
		if newSum > max {
			max = newSum
		}
	}
	fmt.Println(max)
}
