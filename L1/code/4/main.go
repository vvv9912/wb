package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/
func main() {
	nworker := 3
	//arr := []int{2, 4, 6, 8, 10}
	//arr = "sssss"
	interfaceCh := make(chan interface{}, nworker)
	//
	go func() {
		for {
			interfaceCh <- rand.Int()
		}
	}()
	for i := 0; i < nworker; i++ {
		go worker(interfaceCh)
	}
	//завершение программы
	с := make(chan os.Signal)
	signal.Notify(с, os.Interrupt)
	<-с
	close(interfaceCh) //закрываем канал

}

func worker(interfaceCh chan interface{}) {
	val := <-interfaceCh
	os.Stdout.WriteString(fmt.Sprintf("%v\n", val))
}
