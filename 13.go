// 13. Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	var number_one int = 15
	var number_two int = 25

	fmt.Printf("First number: %d\nSecond number: %d\n", number_one, number_two)
	// проверяем место в памяти
	fmt.Printf("%p\n", &number_one)
	fmt.Printf("%p\n", &number_two)
	// просто переназначаем обе переменные таким образом
	number_one, number_two = number_two, number_one
	fmt.Printf("First number after a change: %d\nSecond number after a change: %d\n", number_one, number_two)
	// проверяем место в памяти после изменения значений
	fmt.Printf("%p\n", &number_one)
	fmt.Printf("%p\n", &number_two)
}
