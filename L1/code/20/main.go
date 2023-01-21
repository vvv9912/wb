package main

import "strings"

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
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
