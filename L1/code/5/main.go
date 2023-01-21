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
func main() {
	interfaceCh := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { //таймер завершение программы
		close(interfaceCh) //закрываем канал
		os.Exit(1)
	})
	go worker(interfaceCh)
	for {
		interfaceCh <- rand.Int()
	}

}

func worker(interfaceCh chan interface{}) {
	for {
		val := <-interfaceCh
		_, err := os.Stdout.WriteString(fmt.Sprintf("%v\n", val))
		if err != nil {
			return
		}
	}
}
