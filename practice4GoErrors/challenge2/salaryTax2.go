// Haz lo mismo que en el ejercicio anterior pero reformulando el
// código para que, en reemplazo de “Error()”,  se implemente
// “errors.New()”.

package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 160000

	if salary < 150000 {
		fmt.Println(errors.New(" el salario no es suficiente para pagar impuesto"))
		return
	}
	fmt.Println(" Debe pagar impuesto")
}
