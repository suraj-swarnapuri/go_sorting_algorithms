package comparisonsort

import "golang.org/x/exp/constraints"

func mergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge[T constraints.Ordered](a []T, b []T) []T {
	final := []T{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func MergeSort[T constraints.Ordered](inputArray []T) []T {
	// make a copy of slice
	arr := make([]T, len(inputArray))
	copy(arr, inputArray)
	return mergeSort(arr)
}
