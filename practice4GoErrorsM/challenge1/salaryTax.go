// 1. En tu función “main”, define una variable llamada
// “salary” y asignarle un valor de tipo “int”.
// 2. Crea un error personalizado con un struct que implemente
// “Error()” con el mensaje “error: el salario ingresado
// no alcanza el mínimo imponible" y lánzalo en caso de que
// “salary” sea menor a 150.000. Caso contrario, imprime por
// consola el mensaje “Debe pagar impuesto”.

package main

import "fmt"

func main() {
	err := CustomErrorTest(140000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error: el salario ingresado no alcanza el mínimo imponible")
}

func CustomErrorTest(salary int) error {
	if salary < 150000 {
		return &CustomError{
			msg: "error ejecutado",
		}
	}
	return nil
}
