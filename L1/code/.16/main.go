package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/
//todo
func main() {
	arr := []int{1, 2, 34, 152, 100, 5, 32, 100, 0}
	quickSort(arr)
}
func quickSort(a []int) {
	//
	for i := 0; i < len(a)-1; i++ {
		ai := a[i]
		al := a[len(a)-1]
		_ = ai
		_ = al
		if a[i] > a[len(a)-1-i] {
			a[i], a[len(a)-1] = a[len(a)-1], a[i]
		}
	}
	println(a)
	b := a
	_ = b
}
