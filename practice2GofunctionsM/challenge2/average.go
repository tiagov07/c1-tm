// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
// y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo.

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(average(2, 5, 3))
}

func average(ratings ...int) (float32, error) {
	var result float32
	for _, rating := range ratings {
		if rating < 0 {
			return 0, errors.New("function don't allow negative values")
		} else {
			result += float32(rating)
		}
	}
	return result / float32(len(ratings)), nil
}
