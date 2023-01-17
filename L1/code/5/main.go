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

	go func() {
		for {

			interfaceCh <- rand.Int()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			go worker(interfaceCh)
		}
	}()

	//завершение программы
	time.AfterFunc(time.Second, func() {
		fmt.Printf("Таймер выполнен")
		close(interfaceCh) //закрываем канал
		os.Exit(1)
	})

	//close(interfaceCh) //закрываем канал

}

func worker(interfaceCh chan interface{}) {
	val := <-interfaceCh
	_, err := os.Stdout.WriteString(fmt.Sprintf("%v\n", val))
	if err != nil {
		return
	}
}
