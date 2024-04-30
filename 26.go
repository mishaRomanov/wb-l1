// 26. Разработать программу, которая проверяет, что все символы
// в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {
	// используем пакет стрингс для перевода строки в нижний регистр
	str = strings.ToLower(str)
	// создаем мапу для проверки уникальности
	map_check := make(map[rune]struct{})
	for _, char := range str {
		// если в мапе уже существует этот символ, возвращаем false.
		// в противном случае записываем в false
		if _, exists := map_check[char]; exists {
			return false
		} else {
			map_check[char] = struct{}{}
		}
	}
	// если за весь цикл ни разу не сработал фолс, возвращаем тру
	return true
}

func main() {
	fmt.Println(checkUnique("abcdefgh"))
	fmt.Println(checkUnique("abc"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("abcaBd"))
	fmt.Println(checkUnique("abracadabra"))
	fmt.Println(checkUnique("abcabc"))
}
