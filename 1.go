// 1. Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от
// родительской структуры Human (аналог наследования)
package main

import "fmt"

// Структура human (родительская)
type Human struct {
	Name string
	Age  int
}

// метод для примера
func (h Human) PersonData() string {
	return fmt.Sprintf("%s is %d years old.\n", h.Name, h.Age)
}

// структура action, в которую мы встраиваем родительскую структуру хьюман
// после этого у структуры action будет доступ к полям и данным родительской структуры
// примеры будут ниже
type Action struct {
	Desc string
	Human
}

func (a Action) DoStuff() {
	fmt.Printf("I'm doing %s\n", a.Desc)
}

// тут у нас есть доступ к полям встроенной структуры
func (a Action) Act() {
	fmt.Printf("%s is doing %s\n", a.Name, a.Desc)
}

// а тут мы даже можем получить доступ к встроенным методам. собственно, то что было нужно.
func (a Action) FullInfo() {
	fmt.Printf("%sHe is doing %s\n", a.PersonData(), a.Desc)
}

func main() {
	// создаем родительскую структуру
	hmn := Human{
		"Misha",
		24,
	}
	hmn.PersonData()
	// создаем структура, которая наследует методы родительской структуры благодаря композиции
	actn := Action{
		"Coding",
		hmn,
	}
	// пример использования
	actn.PersonData()
	actn.DoStuff()
	actn.Act()
	actn.FullInfo()

	// I'm doing Coding
	// Misha is doing Coding
	// Misha is 24 years old.
	// He is doing Coding
}
