// 16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка
package main

import "fmt"

// сложность алгоритма - 0(n log n)
func quicksort(list []int) []int {
	// если длина массива меньше 2, то он отсортирован и его можно вернуть
	if len(list) < 2 {
		return list
	}
	// создаем левый и правый массивы - массивы для чисел меньше и больше опорного элемента
	left := []int{}
	right := []int{}

	// выбираем опорный элемент - проще всего сделать это с первым элементом
	pivot := list[0]
	// проходимся по элементам массива
	for _, num := range list[1:] {
		// если опорный элемент больше чем число, то добавляем в массив слева
		if pivot > num {
			left = append(left, num)
			// в противном случае число больше, добавляем в массив справа
		} else {
			right = append(right, num)
		}
	}

	// на данном этапе запускается стэк вызовов, то есть он начинает рекурсивно вызывать функцию сортировки до тех пор,
	// пока массив не будет равен одному элементу, то есть будет отсортирован.

	// левый массив (то есть числа которые меньше чем опорный элемент)
	// делаем отсорированным и добавляем туда в конец опорный элемент
	left = append(quicksort(left), pivot)
	// сортируем правый массив (значения, которые больше опорного элемента)
	right = quicksort(right)

	// возвращаем собранные вместе левый и правый массивы
	return append(left, right...)
}

// тестируем работоспособность в мейн функции
func main() {
	initArray := []int{55, 2, 1, 65436, 33, 55, 17, 12, 9, 5, 0, 4}
	fmt.Println("Starting quicksort test with array: ", initArray)
	fmt.Println("After quicksort we got: ", quicksort(initArray))
}
