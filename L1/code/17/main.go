package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	arr := []int{1, 2, 34, 6, 100, 5, 32, 100, 0}
	n := 12
	fmt.Println(arr)

	sort.Ints(arr)
	fmt.Println(arr)
	a := sort.SearchInts(arr, n) //работает только после сортировки, если не найдено, возвращает в каком диапазоне

	if n == arr[a] {
		fmt.Println(a)
	} else {
		fmt.Println("Число не найдено")
	}

}
