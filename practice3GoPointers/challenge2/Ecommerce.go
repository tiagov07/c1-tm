// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
// productos a los usuarios. Para ello requieren que tanto los usuarios como los productos
// tengan la misma dirección de memoria en el main del programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto
// al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.
package main

import "fmt"

func main() {
	product1 := newProduct("balon", 140000)
	product2 := product{"guantes", 90000, 1}

	fmt.Printf(" first product created %+v \n", product1)
	products := []product{product1}
	newUser := User{"emilio", "casas", "emi@gmail.com", products}
	fmt.Printf(" user created %+v \n", newUser)

	newUser.addProduct("emilio", product2, 2)
	fmt.Printf(" product added %+v \n", newUser)

	newUser.delProducts("emilio")
	fmt.Printf(" products deleted %+v \n", newUser)

}

type User struct {
	Name     string
	LastName string
	Email    string
	Products []product
}

type product struct {
	ProductName string
	price       int
	quantity    int
}

func newProduct(name string, price int) product {
	newProduct := product{ProductName: name, price: price, quantity: 1}
	return newProduct
}

func (u *User) addProduct(Name string, product product, quantity int) {
	if u.Name == Name {
		for i := 0; i < quantity; i++ {
			product.quantity += 1
		}
		u.Products = append(u.Products, product)
	}
	fmt.Println(u)
}

func (u *User) delProducts(Name string) {
	if u.Name == Name {
		u.Products = nil
	}
	fmt.Println(u)
}
