// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
// Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
// haber muchos más animales que refugiar.

// perro necesitan 10 kg de alimento
// gato 5 kg
// Hamster 250 gramos.
// Tarántula 150 gramos.

// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
// especificado y que retorne una función y un mensaje (en caso que no exista el animal)
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del
// tipo de animal especificado.

package main

import (
	"errors"
	"fmt"
)

func main() {
	oper := Animal("dog")
	if oper != nil {
		fmt.Println(oper(2))
	} else {
		fmt.Println(errors.New("animal is not found"))
	}
}

func dogFood(quantity int) float32 {
	return 10 * float32(quantity)
}

func catFood(quantity int) float32 {
	return 5 * float32(quantity)
}

func hamsterFood(quantity int) float32 {
	return 0.25 * float32(quantity)
}

func tarantulaFood(quantity int) float32 {
	return 0.15 * float32(quantity)
}

func Animal(animal string) func(quantity int) float32 {
	switch animal {
	case "dog":
		return dogFood
	case "cat":
		return catFood
	case "hamster":
		return hamsterFood
	case "Tarantula":
		return tarantulaFood
	default:
		return nil
	}
}
