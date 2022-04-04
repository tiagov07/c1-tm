// Se debe implementar la funcionalidad para crear la entidad. pasa eso
// se deben seguir los
// siguientes pasos:
// 1. Crea un endpoint mediante POST el cual reciba la entidad.
// 2. Se debe tener un array de la entidad en memoria (a nivel global),
// en el cual se
// deberán ir guardando todas las peticiones que se vayan realizando.
// 3. Al momento de realizar la petición se debe generar el ID. Para
// generar el ID se debe
// buscar el ID del último registro generado, incrementarlo en 1 y
// asignarlo a nuestro
// nuevo registro (sin tener una variable de último ID a nivel global).
package main

import (
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
var id int

func createProduct(ctxt *gin.Context) {
	token := ctxt.GetHeader("token")
	if token != "123456" {
		ctxt.JSON(401, gin.H{
			"error": "token invalido",
		})
	}
	var req producto
	id++
	req.Id = id
	if error := ctxt.ShouldBindJSON(&req); error != nil {
		ctxt.JSON(400, gin.H{
			"error": error.Error(),
		})
	}

	if req.Nombre == "" {
		ctxt.JSON(400, gin.H{
			"message": "Nombre requerido",
		})
	}
	if req.Color == "" {
		ctxt.JSON(400, gin.H{
			"message": "Color requerido",
		})
	}
	if req.Precio == 0 {
		ctxt.JSON(400, gin.H{
			"message": "Precio requerido",
		})
	}
	if req.Stock == 0 {
		ctxt.JSON(400, gin.H{
			"message": "Stock requerido",
		})

	}
	if req.Publicado == false {
		ctxt.JSON(400, gin.H{
			"message": "Publicado requerido",
		})
	}
	if req.Fecha_creacion == "" {

		ctxt.JSON(400, gin.H{
			"message": "Fecha_creacion requerida",
		})
	}
	if req.Codigo == "" {
		ctxt.JSON(400, gin.H{
			"message": "Codigo requerido",
		})

	}

	productos = append(productos, req)

	ctxt.JSON(200, productos)
	return
}

func valdite(ctxt *gin.Context) {

}

func main() {

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Team 2",
		})
	})

	router.POST("/productos", createProduct)

	router.Run()
}
