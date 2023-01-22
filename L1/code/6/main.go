package main

import (
	"fmt"
	"sync"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
func main() {
	//с помощью канала
	//boolCh := make(chan bool)
	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go worker(boolCh, &wg)
	//boolCh <- true
	//С помощью указателя
	var wg sync.WaitGroup
	var work bool
	//work = false
	wg.Add(1)
	go worker2(&work, &wg)
	//work = true
	wg.Wait()
}

func worker(boolCh chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-boolCh:
			wg.Done()
			fmt.Println("Закрытие горутины")
			return
		default:
			fmt.Println("Горутина работает")
		}
	}
}
func worker2(boolW *bool, wg *sync.WaitGroup) {
	for {
		if boolW != nil {
			if *boolW {
				wg.Done()
				fmt.Println("Закрытие горутины")
				return
			}
		}
		fmt.Println("Горутина работает")
	}
}
