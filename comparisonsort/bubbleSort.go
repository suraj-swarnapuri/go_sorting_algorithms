package comparisonsort

func BubbleSort(inputArray []int) []int {
	// make a copy of slice
	arr := make([]int, len(inputArray))
	copy(arr, inputArray)
	// sort array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				//swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// return sorted array
	return arr
}
