package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id             int
	Nombre         string
	Color          string
	Precio         float64
	Stock          int
	Codigo         string
	Publicado      bool
	fecha_creacion string
}

func GetAll(c *gin.Context) {
	var p []producto
	x, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal("error")
	}
	err = json.Unmarshal(x, &p)
	if err != nil {
		log.Fatal("Otro error")
	}
	c.JSON(200, gin.H{"productos": p})
}

func main() {

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Team 2",
		})
	})

	router.GET("/productos", GetAll)

	router.Run()
}
