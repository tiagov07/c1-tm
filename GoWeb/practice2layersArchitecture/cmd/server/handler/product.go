package handler

import (
	"github/tiagov07/c1-tm/GoWeb/practice2layersArchitecture/internal/products"
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

type ProductHandler struct {
	service products.Service
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (p *ProductHandler) GetAll() gin.HandlerFunc {
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

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *ProductHandler) Store() gin.HandlerFunc {
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

func NewProduct(s products.Service) *ProductHandler {
	return &ProductHandler{service: s}
}
