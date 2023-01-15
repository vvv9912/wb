package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	massiv := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(massiv))
	for i := range massiv {
		go func(i int) {
			os.Stdout.WriteString(fmt.Sprintf("%d\n", massiv[i]*massiv[i]))
			wg.Done()
		}(i)

	}
	wg.Wait()

}
