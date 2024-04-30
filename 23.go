// 23. Удалить i-ый элемент из слайса.
package main

import "fmt"

func removeByIndex(list []int, index int) []int {
	// функция аппенд возвращает слайс, который состоит
	// из среза прошлого слайса ДО индекса числа и ПОСЛЕ индекса числа
	// синтакс '...' после второго среза означает, что там находится больше чем 1 элемент
	return append(list[:index], list[index+1:]...)
}

func main() {
	var slice []int
	for n := range 10 {
		slice = append(slice, n)
	}
	fmt.Println("Оригинальный слайс:", slice)
	slice = removeByIndex(slice, 3)
	fmt.Println("Слайс после удаления числа по индексу 3:", slice)
}
