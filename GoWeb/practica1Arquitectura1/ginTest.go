package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id             int     `form:"Id" json:"Id"`
	Nombre         string  `form:"Nombre" json:"Nombre"`
	Color          string  `form:"Color" json:"Color"`
	Precio         float64 `form:"Precio" json:"Precio"`
	Stock          int     `form:"Stock" json:"Stock"`
	Codigo         string  `form:"Codigo" json:"Codigo"`
	Publicado      bool    `form:"Publicado" json:"Publicado"`
	Fecha_creacion string  `form:"Fecha_creacion" json:"Fecha_creacion"`
}

var productos []producto

func GetAll(c *gin.Context) {
	var jsonData, err = os.ReadFile("products.json")
	if err != nil {
		log.Fatal("error")
	}
	err = json.Unmarshal(jsonData, &productos)
	if err != nil {
		log.Fatal("Otro error")
	}
	c.JSON(200, gin.H{"productos": productos})
}

/*
Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint,
el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.
*/

func getProduct(ctxt *gin.Context) {
	var productList []producto
	jsonData, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &productList)
	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(ctxt.Query("Id"))

	if err != nil {
		fmt.Println(err)
	}

	var productFilter []producto
	for _, product := range productList {
		fmt.Println(product)
		if product.Id == id {
			productFilter = append(productFilter, product)
		}
	}
	ctxt.JSON(200, productFilter)
}

/*
Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del array
de la temática. Utilizando path parameters el endpoint debería ser /temática/:id
(recuerda que siempre tiene que ser en plural la temática).
Una vez recibido el id devuelve la posición correspondiente.
Genera una nueva ruta.
Genera un handler para la ruta creada.
Dentro del handler busca el item que necesitas.
Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/

func getProductByParam(ctxt *gin.Context) {
	var productList []producto
	jsonData, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &productList)
	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(ctxt.Param("Id"))

	if err != nil {
		fmt.Println(err)
	}

	var productFilter []producto
	for _, product := range productList {
		fmt.Println(product)
		if product.Id == id {
			productFilter = append(productFilter, product)
		}
	}
	ctxt.JSON(200, productFilter)
}

func main() {

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Team 2",
		})
	})

	router.GET("/productos", GetAll)
	router.GET("/product", getProduct)
	router.GET("/productByParam/:id", getProductByParam)

	router.Run()
}
