// 11. Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func SetIntersection(set_one, set_two map[string]struct{}) map[string]struct{} {
	// делаем мапу для пересечения множеств
	intersection := make(map[string]struct{})
	for k := range set_one {
		// если элемент одного множества существует в другом,
		// добавляем в пересечение
		if _, exists := set_two[k]; exists {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func main() {
	// первое множество
	set_one := map[string]struct{}{
		"flower": {},
		"cat":    {},
		"tree":   {},
		"chair":  {},
		"dog":    {},
		"mouse":  {},
		"bird":   {},
	}
	// второе множество
	set_two := map[string]struct{}{
		"cat":   {},
		"dog":   {},
		"mouse": {},
		"bird":  {},
		"fish":  {},
		"bug":   {},
	}
	fmt.Println(SetIntersection(set_one, set_two))
}
