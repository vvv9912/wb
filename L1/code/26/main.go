package main

import "strings"

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
*/

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	b := unique(s1)
	println(s1, " = ", b)
	b = unique(s2)
	println(s2, " = ", b)
	b = unique(s3)
	println(s2, " = ", b)
}
func unique(s string) bool {
	//переведем сразу в нижний регистр
	runs := []rune(strings.ToLower(s))
	for m, v := range runs {
		for i := m + 1; i < len(runs); i++ {
			if v == runs[i] {
				return false
			}
		}
	}
	return true
}
