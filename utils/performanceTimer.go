package utils

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

func PerformanceTimer[T constraints.Ordered](funcName string, f func(arr []T) []T, arr []T, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	f(arr)
	end := time.Now()
	fmt.Printf("%s Time taken: %s\n", funcName, end.Sub(start))
}

func PerformanceTimerStable(funcName string, f func(arr []int, max int) []int, arr []int, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	f(arr, max)
	end := time.Now()
	fmt.Printf("%s Time taken: %s\n", funcName, end.Sub(start))
}
