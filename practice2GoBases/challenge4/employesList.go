// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
// Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

// Por otro lado también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func main() {
	employe := ""
	fmt.Println(" type the employe's name you want to work with please : ")
	fmt.Scanf("%s", &employe)

	employee := employees[employe]
	if employee != 0 {
		fmt.Printf("the age of the employee is : %v years old \n", employee)
	} else {
		fmt.Println(" be sure you type the name in the correct way")
	}
	m := 0
	for _, p := range employees {
		if p > 21 {
			m++
		}
	}
	fmt.Println("employees with more than 21 years old are : ", m)

}
