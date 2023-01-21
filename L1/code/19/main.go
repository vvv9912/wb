package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	arr := "главрыба"
	println(arr)
	rev(&arr)
	println(arr)
}
func rev(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}
