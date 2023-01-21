package main

import (
	"fmt"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/
var justString string

func createHugeString(len int) string {
	var v string
	for i := 0; i < len; i++ {
		v += string(195) // Ã
		//если взять разные символы из utf8
	}
	return v
}
func someFunc() {
	v := createHugeString(1 << 10)

	justString = v[:100] // берем 100
	fmt.Println(justString[0], string(justString[0]))
	fmt.Println(justString[1], string(justString[1]))
	fmt.Println(justString)
	justString2 := string([]rune(v))
	fmt.Println(justString2)
}

func main() {
	someFunc()
}
