// Una empresa nacional se encarga de realizar venta de productos,
// servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular
// el precio total de Productos, Servicios y Mantenimientos. Debido a la
// fuerte demanda y para optimizar la velocidad requieren que el cálculo
// de la sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio
// total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio
// total (precio * media hora trabajada, si no llega a trabajar 30 minutos
// se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el
// precio total.

// Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por
// pantalla el monto final (sumando el total de los 3).

package main

import "fmt"

func main() {
	product1 := Product{"auto", 30000000, 2}
	// Product2 := Product{"lampara", 15000, 1}
	// Product3 := Product{"Ambientador", 300, 4}
	service1 := Service{"lavado", 3000, 2}
	// service2 := Service{"polichado", 5000, 1}
	maintenance1 := Maintenance{"cambio Aceite", 50000}

	fmt.Println("here we go")
	chanel := make(chan int) // in order to call the main function after
	// the go routine functions
	products := []Product{product1}
	go product1.productCost(products, chanel)
	services := []Service{service1}
	go service1.servicesCost(services, chanel)
	maintenances := []Maintenance{maintenance1}
	go maintenance1.maintenanceCost(maintenances, chanel)
	totalCost := <-chanel + <-chanel + <-chanel
	fmt.Println(totalCost)
}

type Product struct {
	Name     string
	Price    int
	Quantity int
}

type Service struct {
	Name          string
	Price         int
	MinutesOfWork int
}

type Maintenance struct {
	Name  string
	Price int
}

func (p *Product) productCost(pts []Product, chanel chan int) {
	totalCost := 0
	for i := 0; i < len(pts); i++ {
		totalCost += pts[i].Price * pts[i].Quantity
	}
	fmt.Printf(" products cost  ")
	chanel <- totalCost

}

func (s *Service) servicesCost(srs []Service, chanel chan int) {
	totalCost := 0
	for i := 0; i < len(srs); i++ {
		totalCost += srs[i].Price * srs[i].MinutesOfWork
	}
	fmt.Printf(" service cost ")
	chanel <- totalCost
}

func (m *Maintenance) maintenanceCost(srs []Maintenance, chanel chan int) {
	totalCost := 0
	for i := 0; i < len(srs); i++ {
		totalCost += srs[i].Price
	}
	fmt.Printf(" maintenances cost ")
	chanel <- totalCost
}
