// Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
// Las empresas tienen 3 tipos de productos:
// Pequeño, Mediano y Grande. (Se espera que sean muchos más)
// Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

// Sus costos adicionales son:
// Pequeño: El costo del producto (sin costo adicional)
// Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
// Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

// Requerimientos:
// Crear una estructura “tienda” que guarde una lista de productos.
// Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
// Crear una interface “Producto” que tenga el método “CalcularCosto”
// Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
// Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
// Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
// Interface Producto:
// El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
// Interface Ecommerce:
//  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
//  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

package main

import "fmt"

func main() {
	product1 := newProduct("little", "pizza", 1000)
	product2 := newProduct("medium", "burguer", 4000)
	product3 := newProduct("big", "barril", 15000)

	store1 := newStore()
	store1.Add(product1)
	store1.Add(product2)
	store1.Add(product3)

	fmt.Println(store1.Total())
}

type Product interface {
	getCost() float64
}

type Ecommerce interface {
	Total() float64
	Add(p product)
}

type store struct {
	products []product
}

func (s *store) Add(p product) {
	s.products = append(s.products, p)
}

func (s store) Total() float64 {
	var total float64
	for _, p := range s.products {
		total += p.Price + p.getCost()
	}
	return total
}

type product struct {
	Type  string
	Name  string
	Price float64
}

func newProduct(Type string, Name string, Price float64) product {
	return product{Type, Name, Price}
}

func newStore() Ecommerce {
	return &store{[]product{}}
}

func (p product) getCost() float64 {
	if p.Type == "little" {
		return 0
	} else if p.Type == "medium" {
		return p.Price * 0.03
	} else if p.Type == "big" {
		return p.Price*0.06 + 2500
	} else {
		return 0
	}

}
