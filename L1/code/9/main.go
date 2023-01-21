package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/
func main() {
	chInt1 := make(chan int) //пишем х
	chInt2 := make(chan int) // пишем x*2
	var wg sync.WaitGroup
	wg.Add(1) //добавляем для не завершения программы

	go respch1(chInt1, &wg)
	go respch2(chInt1, chInt2)
	go answ(chInt2, &wg)

	wg.Wait()

}
func respch1(chInt1 chan int, wg *sync.WaitGroup) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	wg.Add(len(arr) - 1)
	for i := range arr {
		chInt1 <- arr[i]
		fmt.Printf("Отправлено канал 1\n")
	}
}
func respch2(chInt1 chan int, chInt2 chan int) {
	for v := range chInt1 {
		fmt.Println(v)
		chInt2 <- v * 2
		fmt.Printf("Обработано в канале 2\n")
	}
}
func answ(chInt2 chan int, wg *sync.WaitGroup) {
	for v := range chInt2 {
		fmt.Printf("Ответ из канала 2: %d \n", v)
		wg.Done()
	}
}
