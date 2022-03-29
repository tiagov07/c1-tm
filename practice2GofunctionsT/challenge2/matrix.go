// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix (vector de vectores)que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre rows)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y
// cuál es el valor máximo.

package main

import (
	"errors"
	"fmt"
	"log"
)

type Matriz struct {
	Elements [][]float64
	Height   int
	Width    int
}

func (m *Matriz) Set(rows ...[]float64) error {
	if m.Height < len(rows) {
		return errors.New(" the matriz height is less than the quantity of inputs")
	}

	for i, val := range rows {
		if len(val) > m.Width {
			return errors.New(" the matriz width is less than the quantity of inputs")
		}
		m.Elements[i] = val
	}
	return nil
}

func (m Matriz) Print() {
	for _, row := range m.Elements {
		for _, item := range row {
			fmt.Printf("%.2f\t", item)
		}
		fmt.Println()
	}
}

func newMatriz(height, width int) Matriz {
	m := Matriz{Height: height, Width: width}

	for i := 0; i < height; i++ {
		m.Elements = append(m.Elements, make([]float64, width))
	}
	return m
}
func main() {
	m := newMatriz(3, 4)
	fmt.Printf("%+v\n", m)
	Elements := make([][]float64, 3)
	Elements[0] = []float64{4, 7.6, 8, -1}
	Elements[1] = []float64{5, 7.6, 8, 4}
	Elements[2] = []float64{4, 8, 8, 2}

	err := m.Set(Elements[0], Elements[1], Elements[2])
	if err != nil {
		log.Fatal(err)
	}
	m.Print()
}
