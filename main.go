package main

import "fmt"

func main() {
	var x, y int
	var znak string

	fmt.Println()
	fmt.Println("Введите данные, cоблюдая пробелы, например 2 + 2: ")
	fmt.Scan(&x, &znak, &y)

	switch znak {
	case "+":
		fmt.Printf("= %1d", Plus(x, y))
	case "-":
		fmt.Printf(" = %1d", Minus(x, y))
	case "*":
		fmt.Printf(" = %1d", Mult(x, y))
	case "/":
		fmt.Printf(" = %d", Div(x, y))
	default:
		fmt.Println("Ошибка,введите выражение, еще раз:     ")
		fmt.Scan(&x, &znak, &y)
	}
	fmt.Println("\n \n \n Вычисления завершены, можно закрыть программу")
}

func Plus(x, y int) int { return x + y }

func Minus(x, y int) int { return x - y }

func Mult(x, y int) int { return x * y }

func Div(x, y int) int {
	if y == 0 {
		fmt.Printf("Divided by %d \n", y)
	}
	return (x / y)
}
