package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/
//todo
func main() {
	interfaceCh := make(chan interface{})

	time.AfterFunc(10*time.Second, func() {
		fmt.Printf("Таймер выполнен")
		close(interfaceCh) //закрываем канал
		os.Exit(1)
	})
	go worker(interfaceCh)
	for {
		interfaceCh <- rand.Int()
		fmt.Printf("Отправлено\n")
		time.Sleep(time.Second)
	}

	//close(interfaceCh) //закрываем канал

}

func worker(interfaceCh chan interface{}) {
	//var wg sync.WaitGroup
	for {
		val := <-interfaceCh
		_, err := os.Stdout.WriteString(fmt.Sprintf("%v\n", val))
		if err != nil {
			return
		}
	}
	//wg.Wait()
}
