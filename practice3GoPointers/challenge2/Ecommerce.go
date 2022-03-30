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

func main() {
	newproduct := product.newProduct("balon", 140000)
}

type User struct {
	Name     string
	LastName string
	Email    string
	Products []product
}

type product struct {
	ProductName string
	price       uint32
	quantity    uint16
}

func (p product) newProduct(name string, price uint32) product {
	newProduct := product{p.ProductName: name, p.price: price, p.quantity: 1}
	return newProduct
}

func (u *User) addProduct(Name string, product product, quantity uint16) {
	u.Name = Name
	u.Products = append(product)
}

func (u *User) delProducts(Name string) {
	u.Products = (product)
}
