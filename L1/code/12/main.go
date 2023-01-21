package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств
*/
type t_Struct struct {
	a int
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}

	m := map[string]struct{}{}
	for i := range a {
		m[a[i]] = struct{}{}
	}

	for k, _ := range m {
		fmt.Println("Множетсво", k)
	}

}
