// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
// depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva
// el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17%
// del sueldo y si gana más de $150.000 se le descontará además un 10%.

package main

import "fmt"

func tax(salary float32) float32 {
	if salary > 50000 && salary < 150000 {
		salary = salary * 0.17
		fmt.Printf("the salary tax is : %v dolars \n", salary)
	} else {
		salary = salary * 0.27
		fmt.Printf("the salary tax is : %v dolars \n", salary)
	}
	return 0
}

func main() {

	var salary float32 = 1000000

	tax(salary)

}
