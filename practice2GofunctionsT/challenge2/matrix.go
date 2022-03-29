// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y
// cuál es el valor máximo.

package main

type Matrix struct {
	Elements   float32
	Dimensión  float32
	Longitud   float32
	Cuadratica float32
	MaxValue   float32
}

func Set(m Matrix) float32 {
	m := Matrix{Elements: }
}

func Print(m Matrix) string {
	fmt.Printf("Elements:\t%d\nDimension:\t%d\nLongitud:\t%d\nCuadratica:\t%d\nMaxValue:\t%d\n\n", m.Elements, m.Dimensión, m.Longitud, m.Cuadratica, m.MaxValue)

}
