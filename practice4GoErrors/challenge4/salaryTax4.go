// Vamos a hacer que nuestro programa sea un poco más complejo y útil.
// Desarrolla las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora como
// argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le deberá
// descontar el 10% en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un
// número negativo, la función debe devolver un error. El mismo deberá indicar
// “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
// Calcular el medio aguinaldo correspondiente al trabajador
// Fórmula de cálculo de aguinaldo:
// [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
// La función que realice el cálculo deberá retornar más de un valor, incluyendo
// un error en caso de que se ingrese un número negativo.

// Desarrolla el código necesario para cumplir con las funcionalidades requeridas,
// utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides
// realizar las validaciones de los retornos de error en tu función “main()”.

package main

import (
	"errors"
	"fmt"
)

func salarioMensual(valorPorHora float64, horas int) (float64, error) {
	if horas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salarioMensual := valorPorHora * float64(horas)
	if salarioMensual >= 150000 {
		salarioMensual -= salarioMensual * 0.1
	}
	return salarioMensual, nil
}
func wrapError(mensaje string, err error) error {
	if err == nil {
		return errors.New(mensaje)
	}
	return fmt.Errorf(mensaje+": %w", err)
}
func validarParametrosAguinaldo(salario float64, meses int) error {
	var err error = nil
	if salario < 0 {
		err = wrapError("El salario no debe ser negativo", err)
	}
	if meses < 0 {
		err = wrapError("Los meses laborados no deben ser menor a cero", err)
	}
	return err
}
func aguinaldo(salario float64, meses int) (float64, error) {
	if err := validarParametrosAguinaldo(salario, meses); err != nil {
		return 0, err
	}
	return (salario / 12) * float64(meses), nil
}
func main() {
	TestSalario()
	TestAguinaldo()
}
func TestSalario() {
	// El numero de horas es menor a 80 (50) por lo que debe mandar un error argumentando lo mismo
	salario, err := salarioMensual(15, 50)
	fmt.Printf("TestSalario1 -> Salario: %f err: %v\n", salario, err)
	// El salario mensual es menor a 150,000 por lo que no debe aplicar el impuesto
	salario, err = salarioMensual(15, 100)
	fmt.Printf("TestSalario2 -> Salario: %f err: %v\n", salario, err)
	// El salario mensual es 150,000 por lo que debe aplicar el impuesto
	salario, err = salarioMensual(1500, 100)
	fmt.Printf("TestSalario3 -> Salario: %f err: %v\n", salario, err)
}
func TestAguinaldo() {
	// El salario y los meses son mayor a cero por lo que debe calcular bien el aguinaldo y no generar error
	agui, err := aguinaldo(100, 6)
	unwrapError := errors.Unwrap(err)
	fmt.Printf("TestAguinaldo1 -> Aguinaldo: %f err: %v unwrap: %v\n", agui, err, unwrapError)
	// El salario y los meses son mayor a cero por lo que debe calcular bien el aguinaldo y no generar error
	// Los meses son 0 por lo que el aguinaldo debe ser 0
	agui, err = aguinaldo(100, 0)
	unwrapError = errors.Unwrap(err)
	fmt.Printf("TestAguinaldo2 -> Aguinaldo: %f err: %v unwrap: %v\n", agui, err, unwrapError)
	// El salario es negativo por lo que debe mandar un error argumentando el mismo
	agui, err = aguinaldo(-5, 0)
	unwrapError = errors.Unwrap(err)
	fmt.Printf("TestAguinaldo3 -> Aguinaldo: %f err: %v unwrap: %v\n", agui, err, unwrapError)
	// El mes es negativo por lo que debe mandar un error argumentado el mismo
	agui, err = aguinaldo(0, -1)
	unwrapError = errors.Unwrap(err)
	fmt.Printf("TestAguinaldo5 -> Aguinaldo: %f err: %v unwrap: %v\n", agui, err, unwrapError)
	// El salario y los meses son negativos, por lo que debe mandar 2 errores anidados mencionando los mismos
	agui, err = aguinaldo(-5, -1)
	unwrapError = errors.Unwrap(err)
	fmt.Printf("TestAguinaldo4 -> Aguinaldo: %f err: %v unwrap: %v\n", agui, err, unwrapError)
}
