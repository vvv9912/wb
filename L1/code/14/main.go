package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	//a := []string{"cat", "cat", "dog", "cat", "tree"}

	typearr(12)
	typearr("a")
	typearr(true)
	typearr(make(chan int))
	typearr(make(chan string))
	typearr(make(chan bool))
	typearr(make(chan interface{}))
}
func typearr(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	case bool:
		fmt.Println("is bool")
	case chan int, chan string, chan bool:
		fmt.Println("is chan")
	default:
		fmt.Println("unknown")
	}

}
