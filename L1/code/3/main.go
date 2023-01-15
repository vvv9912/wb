package main

import (
	"fmt"
	"sync"
)

func main() {
	massiv := []int{2, 4, 6, 8, 10}
	sum := sumsquares(massiv)
	fmt.Println(sum)
}
func sumsquares(massiv []int) int {
	var sum int
	var wg sync.WaitGroup
	//var rw sync.RWMutex
	wg.Add(len(massiv))
	for i := range massiv {
		go func(i int) {
			//rw.Lock()
			sum += massiv[i] * massiv[i]
			wg.Done()
			//defer rw.Unlock()
		}(i)
	}
	wg.Wait()
	return sum
}
