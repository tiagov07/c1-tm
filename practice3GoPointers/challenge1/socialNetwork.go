// Una empresa de redes sociales requiere implementar una estructura usuario con funciones
// que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
// que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para
// las funciones:
// La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
// Y deben implementarse las funciones:
// cambiar nombre: me permite cambiar el nombre y apellido.
// cambiar edad: me permite cambiar la edad.
// cambiar correo: me permite cambiar el correo.
// cambiar contraseña: me permite cambiar la contraseña.
package main

import "fmt"

func main() {
	newUser := Network{"Vicente", "Benavidez", 23, "chente@gmail.com", "1234"}
	fmt.Printf("%+v \n", newUser)

	newUser.updateName("Carlos", "Fernandez")
	fmt.Printf("%+v \n", newUser)

	newUser.updateAge(25)
	fmt.Printf("%+v \n", newUser)

	newUser.updateEmail("chentico@gmail.com")
	fmt.Printf("%+v \n", newUser)

	newUser.updatePassword("keep-faith")
	fmt.Printf("%+v \n", newUser)
}

type Network struct {
	Name     string
	LastName string
	Years    uint16
	Email    string
	Password string
}

func (n *Network) updateName(Name, LastName string) {
	n.Name = Name
	n.LastName = LastName
}

func (n *Network) updateAge(Years uint16) {
	n.Years = Years
}

func (n *Network) updateEmail(Email string) {
	n.Email = Email
}

func (n *Network) updatePassword(Password string) {
	n.Password = Password
}
