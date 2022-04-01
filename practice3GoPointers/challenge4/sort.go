// Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento
//  utilizar para sus servicios.
// Para ellos se requiere instanciar 3 arreglos con valores aleatorios
// desordenados
// un arreglo de números enteros con 100 valores
// un arreglo de números enteros con 1000 valores
// un arreglo de números enteros con 10000 valores

// Para instanciar las variables utilizar rand

// Se debe realizar el ordenamiento de cada una por:
// Ordenamiento por inserción
// Ordenamiento por burbuja
// Ordenamiento por selección

// Una go routine por cada ejecución de ordenamiento
// Debo esperar a que terminen los ordenamientos de 100 números para
// seguir el de 1000 y después el de 10000.
// Por último debo medir el tiempo de cada uno y mostrar en pantalla el
// resultado, para saber qué ordenamiento fue mejor para cada arreglo

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	//Generate a random array of length n
	canal := make(chan int)
	n100 := generateSlice(100)
	start100 := time.Now()
	go insertionSort(n100, canal)
	finish100 := time.Since(start100)
	log.Printf("insertion sort took : %v  \n", finish100)
	n1000 := generateSlice(1000)
	start1000 := time.Now()
	go bubbleSort(n1000, canal)
	finish1000 := time.Since(start1000)
	log.Printf("bubble sort took : %v \n", finish1000)
	n10000 := generateSlice(10000)
	start10000 := time.Now()
	go selectionSort(n10000, canal)
	finish10000 := time.Since(start10000)
	log.Printf("selection sort took : %v \n", finish10000)
	fmt.Printf(" procesos pendientes : %v  \n| %v \n| %v \n", <-canal, <-canal, <-canal)
}

// slice generator
func generateSlice(size int) []int {
	slice := make([]int, size, size) //first size length, second size capacity
	//seed
	rand.Seed(time.Now().UnixNano())
	for i, _ := range slice {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// insertion function
func insertionSort(items []int, c chan int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	c <- 0
}

//bubble sort function
func bubbleSort(items []int, c chan int) {
	for i := len(items); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if items[j-1] > items[j] {
				intermediate := items[j]
				items[j] = items[j-1]
				items[j-1] = intermediate

			}
		}
	}
	c <- 0
}

// selection sort function
func selectionSort(items []int, c chan int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	c <- 0
}
