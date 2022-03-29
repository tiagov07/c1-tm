// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno
// de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
package main

import "fmt"

type Student struct {
	Name     string
	Lastname string
	DNI      int
	Date     string
}

func (s Student) detail() {
	fmt.Printf("Name:\t%s\nLastname:\t%s\nDNI:\t%d\nDate:\t%s\n\n", s.Name, s.Lastname, s.DNI, s.Date)
}

func main() {
	student1 := Student{Name: "Isaura", Lastname: "Rendon", DNI: 303948321, Date: "03-08-2016"}
	student1.detail()
	student2 := Student{Name: "Emilio", Lastname: "Valencia", DNI: 323948421, Date: "10-08-2016"}
	student2.detail()
}
