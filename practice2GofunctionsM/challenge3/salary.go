// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad
// de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por
// mes y la categoría, y que devuelva su salario.
package main

import "fmt"

func main() {
	fmt.Println(salary(120, "A"))
	fmt.Println(salary(120, "B"))
	fmt.Println(salary(120, "C"))
	fmt.Println(salary(120, "T"))
}

func salary(min int, category string) float32 {
	var hours float32 = float32(min) / 60
	switch category {
	case "C":
		return hours * 1000
	case "B":
		return (hours * 1500) * 1.2
	case "A":
		return (hours * 3000) * 1.5
	default:
		return 0
	}
}
