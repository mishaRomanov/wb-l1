// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func Reverse(source string) string {
	// делаем билдер для более эффективного использования памяти
	var final strings.Builder
	// разделяем строку на элементы с помощью сепаратора пробела
	split := strings.Split(source, " ")
	// проходимся по массиву
	for i := len(split) - 1; i >= 0; i-- {
		// если это последний элемент, то пробел не нужен
		if i == 0 {
			final.WriteString(split[i])
			// записываем еще и пробел
		} else {
			final.WriteString(split[i] + " ")
		}
	}
	// возвращаем в виде строки
	return final.String()
}

func main() {
	str := "sun dog flower coffee"
	fmt.Println(Reverse(str))
}
