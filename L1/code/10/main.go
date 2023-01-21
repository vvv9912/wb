package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func main() {
	arr := []float64{-15.9, -10.0, -20, -19.9, -25, -23.3, -20, -7, -5, 7, 5.1, 6.5, 11.1, 10, 20, 29, 30, 31, 15.2, 21.4, 25.5}

	var m map[int][]float64
	m = make(map[int][]float64)

	for i := range arr {
		if arr[i] >= 0 {
			mass := math.Floor(arr[i]/10) * 10 //округление
			a := m[int(mass)]
			a = append(a, arr[i])
			m[int(mass)] = a
		} else {
			mass := math.Ceil(arr[i]/10) * 10 //округление
			a := m[int(mass)]
			a = append(a, arr[i])
			m[int(mass)] = a
		}
	}
	var i int
	for a, m := range m {
		for _, v := range m {
			i++
			fmt.Println("Множетсво", a, "Значение", v)
		}
	}
	fmt.Println(len(arr))
	fmt.Println(i)
}
