package main

import "fmt"

type Human struct {
	a int
}
type Action struct {
	Human
	b int
}

func main() {
	var action Action
	action.b = 1
	action.a = 2
	fmt.Println(action)
}
