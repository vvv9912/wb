package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := sumsquares(arr)
	fmt.Println(sum)
}
func sumsquares(arr []int) int {
	var sum int
	var wg sync.WaitGroup
	var rw sync.RWMutex
	wg.Add(len(arr))
	for i := range arr {
		go func(i int) {
			rw.Lock()
			sum += arr[i] * arr[i]
			wg.Done()
			defer rw.Unlock()
		}(i)
	}
	wg.Wait()
	return sum
}
