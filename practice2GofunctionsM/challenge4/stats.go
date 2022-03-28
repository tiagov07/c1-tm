// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
// calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo
// y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
// (mínimo, máximo o promedio) y que devuelva otra función
// ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N
// de enteros y devuelva el cálculo que se indicó en la función anterior
package main

import "fmt"

func main() {
	oper := calculationType("prom")
	if oper != nil {
		fmt.Println(oper(8, 5, 1))
	}

}

func min(notes ...int) float32 {
	var min int
	for i, note := range notes {
		if i == 0 {
			min = note
		} else if note < min {
			min = note
		}
	}
	return float32(min)
}
func max(notes ...int) float32 {
	var max int
	for i, note := range notes {
		if i == 0 {
			max = note
		} else if note > max {
			max = note
		}
	}
	return float32(max)
}

func prom(notes ...int) float32 {
	var res float32
	for _, note := range notes {
		res += float32(note)
	}
	return res / float32(len(notes))
}
func calculationType(tipo string) func(notes ...int) float32 {
	switch tipo {
	case "minimo":
		return min
	case "maximo":
		return max
	case "promedio":
		return prom
	default:
		fmt.Println("calculation is not defined")
		return nil
	}
}
