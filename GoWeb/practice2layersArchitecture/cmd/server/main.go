package main

import (
	"github/tiagov07/c1-tm/GoWeb/practice2layersArquitecture/cmd/server/handler"
	"github/tiagov07/c1-tm/GoWeb/practice2layersArquitecture/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	entityProduct := handler.NewProduct(service)

	route := gin.Default()

	group := route.Group("/products")

	group.GET("/", entityProduct.GetAll())
	group.POST("/", entityProduct.Store())
	route.Run()

}
