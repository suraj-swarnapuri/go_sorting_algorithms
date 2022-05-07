package main

import (
	"sync"

	"github.com/suraj-swarnapuri/go_sorting_algorithms/comparisonsort"
	"github.com/suraj-swarnapuri/go_sorting_algorithms/stablesort"
	"github.com/suraj-swarnapuri/go_sorting_algorithms/utils"
)

func main() {
	max := 10000
	arr := utils.GenerateRandomArray(max)
	wg := &sync.WaitGroup{}
	wg.Add(4)
	// Test bubble sort performance
	go utils.PerformanceTimer("Bubble Sort:", comparisonsort.BubbleSort[int], arr, wg)
	// Test insertion sort performance
	go utils.PerformanceTimer("Insertion Sort:", comparisonsort.InsertionSort[int], arr, wg)
	// Test merge sort performance
	go utils.PerformanceTimer("Merge Sort:", comparisonsort.MergeSort[int], arr, wg)
	// Test counting sort performance
	go utils.PerformanceTimerStable("Counting Sort:", stablesort.CountingSort, arr, max, wg)

	wg.Wait()
}
