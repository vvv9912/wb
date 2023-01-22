package main

import (
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func Sleep(t int) {

	<-time.After(time.Second * time.Duration(t))
	//time.AfterFunc(time.Second*time.Duration(t), func() {
	//	fmt.Println(t, "Секунд прошло")
	//	return
	//})
}
func main() {
	s := 5
	Sleep(s)

}
