package main

import (
	"fmt"
	"math"
)

type Point struct {
	//если название начинается заглавной буквы — это public-доступ,
	//если со строчной — private.
	x, y float64
}

// конструктор
func NewPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// AB = √(xb - xa)2 + (yb - ya)2
func Distance(a Point, b Point) float64 {
	//чтобы найти дистанцию надо
	//координаты одной точки вычесть координату другой точки.
	dx := a.x - b.x
	dy := a.y - b.y
	//умножаем друг на друга, чтобы возвести в квадрат
	//Sqrt возвращает квадратный корень из x
	return math.Sqrt(float64(dx*dx + dy*dy))
}

// Найти расстояние между точками A(-1, 3) и B(6,2).
func main() {
	A := NewPoint(-1, 3)
	B := NewPoint(6, 2)

	fmt.Println(Distance(A, B))
}
