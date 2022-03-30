// Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
// para que el mensaje de error reciba por parámetro el valor de “salary”
// indicando que no alcanza el mínimo imponible (el mensaje mostrado por
// 	consola deberá decir: “error: el mínimo imponible es de 150.000 y
// 	el salario ingresado es de: [salary]”, siendo [salary] el valor de
// 	tipo int pasado por parámetro).

package main

import "fmt"

func main() {
	salary := 140000
	if salary < 150000 {
		error := fmt.Errorf(" el minimo imponible es de 150000 y el salario ingresado es : %v ", salary)
		fmt.Println("error ocurrido : ", error)
		return
	}
	fmt.Println(" debe pagar impuesto")
}
