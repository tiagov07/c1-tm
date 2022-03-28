// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por
// separado para deletrearla.
// Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
// Luego imprimí cada una de las letras.
package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) int {
	words := strings.Split(s, "")
	m := 0
	for _, word := range words {
		fmt.Println(word)
		m++
	}
	return m
}

func main() {
	var word string
	fmt.Println("let us know the word you want to work with : ")
	fmt.Scanf("%s", &word)
	wordL := wordCount(word)
	fmt.Printf(" the length of the word is : %v \n", wordL)
}
