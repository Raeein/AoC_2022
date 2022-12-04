package sliceHelper

import "strconv"

func MakeIntSlice(stringArray []string) []int {
	intArray := make([]int, len(stringArray))
	for i := range stringArray {
		intArray[i], _ = strconv.Atoi(stringArray[i])
	}
	return intArray
}

func SumIntSlice(slice ...int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
