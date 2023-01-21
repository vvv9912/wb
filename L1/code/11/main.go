package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств
*/
//type t_Struct struct {
//	a int
//}

func main() {
	a := []string{"cat", "dog", "bird", "camel"}
	m1 := map[string]struct{}{
		a[0]:     {},
		a[2]:     {},
		"rabbit": {},
		"cow":    {},
	}

	m2 := map[string]struct{}{}
	for i := range a {
		m2[a[i]] = struct{}{}
	}

	m := map[string]struct{}{}
	//Пересечение двух множеств
	for k1, _ := range m2 {
		for k2, _ := range m1 {
			if k1 == k2 {
				m[k1] = struct{}{}
			}
		}
	}
	//for k1, _ := range m2 {
	//	m[k1] = struct{}{}
	//}
	//for k2, _ := range m1 {
	//	m[k2] = struct{}{}
	//}
	for k, _ := range m {
		fmt.Println("Множетсво", k)
	}
}
