package main

import "fmt"

	
func main() {
	var x int
	var znak string
	var y int
	
	fmt.Println()
	fmt.Println("Введите данные cоблюдая пробелы, например 2 + 2: ")
	fmt.Scan( &x, &znak, &y)	

	switch znak {
		case "+":
	    	fmt.Printf("= %1d", Plus(x,y))
		case "-":
			fmt.Printf(" = %1d", Minus(x,y))
		case "*":
			fmt.Printf(" = %1d", Mult(x,y))
		case "/":
			fmt.Printf(" = %1d", Div(x,y))
		default:		
			fmt.Println("Ошибка,введите выражение, еще раз:     ")
			fmt.Scanln(&x, &znak, &y)	
	}

	fmt.Println("\n \n \n Вычисления завершены, можно закрыть программу" )	
}

func Plus(x, y int )int {return x + y}

func Minus(x, y int )int {return x - y}

func Mult(x, y int )int {return x * y}

func Div(x, y int )int {return x / y}