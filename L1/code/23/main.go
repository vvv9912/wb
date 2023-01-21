package main

import (
	"errors"
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	err := remove(&arr, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arr)
	}

}
func remove(a *[]int, i int) error {
	if i > len(*a)-1 {
		return errors.New("индекс больше длины массива")
	}
	*a = append((*a)[:i], (*a)[(i+1):]...)
	return nil
}
