package handler

import (
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/internal/products"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"name" binding:"required"`
	Type  string  `json:"type" `
	Count int     `json:"count" `
	Price float64 `json:"price"`
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
func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		product, error := c.service.GetAll()
		if error != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "elements not found", error.Error()))
			return
		}
		ctx.JSON(200, product)
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
func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "not authorized", "invalid credentials"))
			return
		}
		var req request
		if error := ctx.ShouldBindJSON(&req); error != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "elements not found", error.Error()))
			return
		}
		product, error := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if error != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "elements not found", error.Error()))
			return
		}
		ctx.JSON(200, product)
	}
}

// Update godoc
// @Summary Update products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /products/{id} [put]
func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "not authorized", "invalid credentials"))
			return
		}
		id, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if error != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id not found", error.Error()))
			return
		}
		var req request

		if error := ctx.ShouldBindJSON(&req); error != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "element not found", error.Error()))
		}

		// test reflect

		// requireFields := []string{"Name", "Type", "Count", "Price"}
		// for _, field := range requireFields {
		// 	_, err := dynamicgetter.GetField(&req, field, false)
		// 	if err != nil {
		// 		ctx.JSON(400, gin.H{"error": fmt.Sprintf(" %s is required", field)})
		// 		return
		// 	}
		// }
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, " name required", error.Error()))
			return
		}
		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, " type required", error.Error()))
			return
		}
		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, " count required", error.Error()))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, " price required", error.Error()))
			return
		}
		product, error := c.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
		if error != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "elements not found", error.Error()))
			return
		}
		ctx.JSON(200, product)
	}
}

// UpdateName godoc
// @Summary Update by name the products
// @Tags Products
// @Description update products by name
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "name to update"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /products/{id} [patch]
func (c *ProductHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "not authorized", "invalid credentials"))
			return
		}

		id, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if error != nil {
			ctx.JSON(400, web.NewResponse(400, nil, " id not found", error.Error()))
			return
		}

		var req request
		if error := ctx.ShouldBindJSON(&req); error != nil {
			ctx.JSON(400, web.NewResponse(400, nil, " request don't allow it", error.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, " Name not found", error.Error()))
			return
		}

		product, error := c.service.UpdateName(int(id), req.Name)
		if error != nil {
			ctx.JSON(400, web.NewResponse(400, nil, " element not found", error.Error()))
			return
		}

		ctx.JSON(200, product)
	}
}

// Delete godoc
// @Summary Delete products
// @Tags Products
// @Description delete products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request false "Product to delete"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /products/{id} [delete]
func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "not authorized", "invalid credentials"))
			return
		}
		id, error := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if error != nil {
			ctx.JSON(400, web.NewResponse(400, nil, " id not found", error.Error()))
			return
		}

		error = c.service.Delete(int(id))
		if error != nil {
			ctx.JSON(404, web.NewResponse(404, nil, " element not deleted ", error.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, "element deleted", "", ""))
	}
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{service: p}
}
