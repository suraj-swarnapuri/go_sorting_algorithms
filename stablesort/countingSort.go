package stablesort

func CountingSort(inputArr []int, max int) []int {
	// copy the input array
	arr := make([]int, len(inputArr))
	copy(arr, inputArr)

	// Counting sort
	outputArr := make([]int, len(arr))
	count := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		j := arr[i]
		count[j]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		j := arr[i]
		outputArr[count[j]-1] = j
		count[j]--
	}
	return outputArr
}
