package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for i := range arr {
		go func(i int) {
			os.Stdout.WriteString(fmt.Sprintf("%d\n", arr[i]*arr[i]))
			wg.Done()
		}(i)

	}
	wg.Wait()

}
