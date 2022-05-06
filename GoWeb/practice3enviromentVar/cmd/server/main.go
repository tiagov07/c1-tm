package main

import (
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/cmd/server/handler"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/docs"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/internal/products"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/db"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	_ = godotenv.Load()
	db := db.StorageDB
	// dbT := store.New(store.FileType, "./transactions.json")
	repoProducts := products.NewRepository(db)
	serviceProducts := products.NewService(repoProducts)
	product := handler.NewProduct(serviceProducts)
	route := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	productGroup := route.Group("/products", handler.TokenAuthMiddleware())
	productGroup.POST("/", product.Store())
	productGroup.GET("/", product.GetAll())
	productGroup.GET("/:name", product.GetByName())
	productGroup.PUT("/:id", product.Update())
	productGroup.PATCH("/:id", product.UpdateName())
	productGroup.DELETE("/:id", product.Delete())
	route.Run(":8080")

}
