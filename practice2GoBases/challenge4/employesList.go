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
	name := ""
	fmt.Println(" type the employe's name you want to work with please : ")
	fmt.Scanf("%s", &name)

	employee := employees[name]
	if employee != 0 {
		fmt.Printf("the %v age is : %v years old \n", name, employee)
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
	newE := ""
	years := 0
	fmt.Println(" If you want to add a new employee to our database type his name please : ")
	fmt.Scanf("%s", &newE)
	fmt.Println("Now type his age : ")
	fmt.Scanf("%d", &years)

	employees[newE] = years
	deleteN := ""
	fmt.Println(" If you want to delete some of the list type his name please : ")
	fmt.Scanf("%s", &deleteN)
	delete(employees, deleteN)

	fmt.Printf(" %v \n ", employees)

}
