// 22. Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// блок для объявления переменных
	var (
		x         string
		y         string
		operation string
		fnum      big.Int
		snum      big.Int
		result    big.Int
	)

	//читаем числа из stdin
	//
	fmt.Printf("Enter your first big number\n")
	fmt.Scanln(&x)

	fmt.Printf("Enter your second big number\n")
	fmt.Scanln(&y)

	fmt.Printf("Enter the operation\n")
	fmt.Scanln(&operation)

	// парсим числа из строки
	fnum.SetString(x, 10)
	snum.SetString(y, 10)

	switch operation {
	case "*":
		fmt.Printf("Result: %d\n", result.Mul(&fnum, &snum))
	case "-":
		fmt.Printf("Result: %d\n", result.Sub(&fnum, &snum))
	case "/":
		if snum.Int64() == 0 {
			fmt.Println("Can't divide by zero")
			return
		}
		fmt.Printf("Result: %d\n", result.Div(&fnum, &snum))
	case "+":
		fmt.Printf("Result: %d\n", result.Add(&fnum, &snum))
	default:
		fmt.Println("Unknown operator type. I only support +,/,* and -")
	}
}
