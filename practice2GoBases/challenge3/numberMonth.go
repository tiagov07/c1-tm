// Realizar una aplicación que contenga una variable con el número del mes.
// Según el número, imprimir el mes que corresponda en texto.
// ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
// Ej: 7, Julio

package main

import "fmt"

func main() {
	month := 0
	fmt.Println(" type the number you want to work with please : ")
	fmt.Scanf("%d", &month)
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("you must type a number between 1 or 12 ")
	}
}
