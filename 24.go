// 24. Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

// В Go инкапсуляция реализуется на уровне пакетов.
// Любой обьъект в рамках пакета мы можем сделать либо импортируемым, то есть доступным всем остальным пакетам,
// либо неимпортируемым, то есть инкапсулировать его в рамках пакета.
// за импорт отвечают не ключевые слова, как в других языках программирования, а просто
// названия переменных или объектов. Для импорта используем названия с большой буквы, для инкапсуляции - с маленькой
type Point struct {
	x float64
	y float64
}

// напишем собственный конструктор - функцию, возвращающую элемент структуры
func NewPoint(x, y float64) *Point {
	return &Point{
		x,
		y,
	}
}

// метод который считает расстояние между двумя точками (формулу нашел в инете простите если она неправильная)
func DistanceBetweenPoints(p, p2 Point) float64 {
	return math.Sqrt((p2.x-p.x)*(p2.x-p.x) + (p2.y-p.y)*(p2.y-p.y))
}

// тестим в мейне
func main() {
	p1 := Point{1, 1}
	p2 := Point{4, 4}
	fmt.Println("Расстояние -", DistanceBetweenPoints(p1, p2))
}
