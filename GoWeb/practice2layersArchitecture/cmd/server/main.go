package main

import (
	"github/tiagov07/c1-tm/GoWeb/practice2layersArquitecture/cmd/server/handler"
	"github/tiagov07/c1-tm/GoWeb/practice2layersArquitecture/internal/products"

	"github.com/gin-gonic/gin"
)

// @title MELI Bootcamp API
// @version 1.0
// @description this API Handle MELI products
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
