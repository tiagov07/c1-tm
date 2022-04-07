package main

import (
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/cmd/server/handler"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/internal/products"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repoProducts := products.NewRepository(db)
	serviceProducts := products.NewService(repoProducts)
	product := handler.NewProduct(serviceProducts)
	route := gin.Default()

	productGroup := route.Group("/products")
	productGroup.POST("/", product.Store())
	productGroup.GET("/", product.GetAll())
	productGroup.PUT("/:id", product.Update())
	productGroup.PATCH("/:id", product.UpdateName())
	productGroup.DELETE("/:id", product.Delete())
	route.Run(":3000")

}
