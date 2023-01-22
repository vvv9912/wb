package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/
//todo
func main() {
	arr := []int{1, 2, 34, 152, 100, 5, 32, 100, 0}
	fmt.Println(arr)
	//sort.Ints(arr)
	quickSortStart(arr)
	fmt.Println(arr)

}
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) []int {
	//low - нижняя граница сортировки
	//high вернхняя
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr

}
func partition(arr []int, low, high int) ([]int, int) {
	//выбираем опорную точку
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
