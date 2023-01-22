package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/
type Point struct {
	x float64
	y float64
}

// Конструктор
func Newpoint(x, y float64) Point {
	return Point{y: y, x: x}
}

// Нахождения расстояния
func Distance(x, y Point) float64 {
	return math.Sqrt((x.x-y.x)*(x.x-y.x) + (x.y-y.y)*(x.y-y.y))
}
func main() {
	p1 := Newpoint(1, 2)
	p2 := Newpoint(5, 7)
	d := Distance(p1, p2)
	fmt.Println(d)
}
