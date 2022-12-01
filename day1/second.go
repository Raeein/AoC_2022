package day1

import (
	"fmt"
	"github.com/Raeein/AoC_2022/sliceHelper"
	"sort"
	"strings"
)

func Solve2(input string) {
	snackGroups := strings.Split(input, "\n\n")
	sums := []int{}
	for _, group := range snackGroups {

		groupCalories := strings.Split(group, "\n")
		groupCaloriesInt := sliceHelper.MakeIntSlice(groupCalories)
		newSum := sliceHelper.SumIntSlice(groupCaloriesInt...)
		sums = append(sums, newSum)
		sort.Ints(sums)
		sort.Slice(sums, func(i, j int) bool {
			return sums[i] > sums[j]
		})
	}
	fmt.Println(sliceHelper.SumIntSlice(sums[:3]...))
}
