package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговоезначение счетчика.
*/
type counter struct {
	n  int
	mu sync.RWMutex
	wg sync.WaitGroup
}

func main() {
	count := counter{}
	iterations := 91
	count.wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			count.mu.Lock()
			count.n++
			count.wg.Done()
			count.mu.Unlock()
		}()
	}
	count.wg.Wait()
	fmt.Println(count.n)
}
