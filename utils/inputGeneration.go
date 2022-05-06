package utils

import "math/rand"

func GenerateRandomArray(max int) []int {
	arr := make([]int, max)
	for i := 0; i < max; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}
