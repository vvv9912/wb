package main

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/
func main() {
	var a int64
	a = 14
	var set bool
	set = false
	println(setbit(a, 5, set))

}

func setbit(a int64, n int, set bool) int64 {
	if n > 64 {
		return 0
	}
	nn := powint(2, n)
	// проверка нулевого бита
	//b := math.Log2(float64(a))
	//_ = b
	if set {
		return a | int64(nn)
	} else {
		return a &^ int64(nn) //и не
	}
}
func powint(a int, n int) int {
	var res int
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res = a
	for i := 2; i <= n; i++ {
		res *= a
	}
	return res
}
