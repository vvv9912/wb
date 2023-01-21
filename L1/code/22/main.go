package main

import (
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

func main() {

	a := big.NewFloat(2 << 21)
	b := big.NewFloat(2 << 22)

	println("a = ", a, "b = ", b)
	//перемножение
	a.Mul(a, b)
	println("перемножение", a)
	//Деление
	a.Div(a, b)
	println("деление", a)
	//сумма
	a.Add(a, b)
	println("сумма", a)
	//вычитание
	a.Sub(a, b)
	println("сумма", a)
}
