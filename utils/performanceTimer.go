package utils

import (
	"fmt"
	"sync"
	"time"
)

func PerformanceTimer(funcName string, f func(arr []int) []int, arr []int, wg *sync.WaitGroup) {
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
