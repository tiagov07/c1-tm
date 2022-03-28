// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas
// reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
// se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará
// interés a los que su sueldo es mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

package main

import (
	"fmt"
	"strconv"
)

func approval(y, s int, j string) string {
	condition, _ := strconv.ParseBool(j)
	if y > 22 && condition && s > 1 {
		salary := 0
		fmt.Println("how much is your salary per year ? ")
		fmt.Scanf("%d", &salary)
		if salary > 100000 {
			fmt.Println(" you apply the credit without interest, congrats !! ")

		} else {
			fmt.Println(" you're gonna have a credit with and interest of 2.0%, congrats !!  ")
		}
	} else {
		fmt.Println(" In this moment we can't provide you a credit, you could try in a couple of years ")
	}
	ans := "the study has been finished"
	return ans
}

func main() {
	years := 0
	var job string
	seniority := 0
	fmt.Println("welcome to the customer study, could you share how old are you please : ")
	fmt.Scanf("%d", &years)
	fmt.Println("Do you have a formal job ?, answer t if your answer is yes or f if not please ")
	fmt.Scanf("%s", &job)
	fmt.Println("how much is your seniority in years ? ")
	fmt.Scanf("%d", &seniority)
	approval(years, seniority, job)
}
