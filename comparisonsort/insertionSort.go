package comparisonsort

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](inputArray []T) []T {
	// make a copy of slice
	arr := make([]T, len(inputArray))
	copy(arr, inputArray)
	// sort array
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}
