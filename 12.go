// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func createSetOfStrings(str []string) map[string]struct{} {
	// множество - это коллекция неупорядоченных уникальных элементов
	// поэтому для его реализации идеально подходит map
	// поскольку в мапе все ключи уникальные
	// в качестве значения по ключу можем использовать пустую структуру
	// потому что она ничего не весит
	set := make(map[string]struct{})
	for i := range str {
		set[str[i]] = struct{}{}
	}
	return set
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(createSetOfStrings(arr))
}
