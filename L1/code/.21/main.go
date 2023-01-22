package main

import "strings"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

func main() {
	arr := "snow dog sun"
	println(arr)
	rev(&arr)
	println(arr)
}
func rev(s *string) {
	arrs := strings.Split(*s, " ") //массив слов
	var strnew string
	for i := range arrs {
		strnew += arrs[len(arrs)-1-i] + " "
	}
	*s = strnew
}
