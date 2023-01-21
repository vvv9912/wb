package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/
func main() {
	var m map[string]int
	m = make(map[string]int)
	buf := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var wg sync.WaitGroup
	wg.Add(len(buf))
	for i := range buf {
		go func(i int) {
			m[fmt.Sprintf("%d", buf[i])] = buf[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
}
