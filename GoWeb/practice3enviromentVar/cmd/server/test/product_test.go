package test

import (
	"bytes"
	"encoding/json"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/cmd/server/handler"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/internal/products"
	"github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, pathDB)
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllProducts(t *testing.T) {

	type product struct {
		ID       int `json:"id"`
		Nombre   int `json:"nombre"`
		Tipo     int `json:"tipo"`
		Cantidad int `json:"cantidad"`
		Precio   int `json:"precio"`
	}
	//create server
	r := createServer("products.json")
	//request Get type
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	var objRes []product

	//indicate to server to get the request
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)

}

func TestSaveProduct(t *testing.T) {
	//create server
	r := createServer("products.json")
	//rute post type
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
		"name":"conejillo de indias", "type":"test funcional","count":3,
		"precio": 80.00
	}`)

	// indicate to get request
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateProduct(t *testing.T) {
	//create server
	r := createServer("products.json")

	//rute put type
	req, rr := createRequestTest(http.MethodPut, "/products/4", `{
			"name":"conejillo de indias", "type":"test funcional Update","count":3,
			"price": 120.00
	}`)

	// indicate to get request
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestPatchProduct(t *testing.T) {
	r := createServer("products.json")

	//rute patch type
	req, rr := createRequestTest(http.MethodPatch, "/products/4", `{
			"name":"conejillo de indias en Patch"
	}`)

	// indicate to get request
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDeleteProduct(t *testing.T) {
	r := createServer("products.json")

	//rute delete type
	req, rr := createRequestTest(http.MethodDelete, "/products/4", "")

	// indicate to get request
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
