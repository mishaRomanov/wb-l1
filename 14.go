// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

// может казаться, что можно использовать так называемый type assertion,
// однако он работает не со всеми типами (например проверить, канал это или нет мы не можем)
// поэтому воспользуемся пакетом reflect, который дает возможность в рантайме проверять тип
func TypeCheck(val interface{}) string {
	if reflect.ValueOf(val).IsNil() {
		return "nil"
	}

	switch reflect.ValueOf(val).Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8:
		return "int"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	}
	return "Unknown type"
}

func main() {
	fmt.Println(TypeCheck("a"))
	fmt.Println(TypeCheck(25))
	fmt.Println(TypeCheck(true))
	fmt.Println(TypeCheck(make(chan int)))
}
