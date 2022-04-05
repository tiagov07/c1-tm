package handler

import (
	"github/tiagov07/c1-tm/GoWeb/practice2layersArquitecture/internal/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	Code       string  `json:"code"`
	Posted     *bool   `json:"posted"`
	PostedDate string  `json:"postedDate"`
}

type Product struct {
	service products.Service
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		p, error := p.service.GetAll()
		if error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		var req request

		if error := ctx.ShouldBindJSON(&req); error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
			return
		}

		product, error := p.service.Store(req.Name, req.Stock, req.Price, req.Code, req.Color, *req.Posted, req.PostedDate)
		if error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": error.Error()})
		}
		ctx.JSON(http.StatusOK, product)
	}
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}
