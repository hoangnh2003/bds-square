package controller

import (
	"bds-square-backend/internal/service"
	"bds-square-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) Create(c *gin.Context) {
	id, code := pc.productService.CreateProduct("Coca Cola", "10000")
	response.SuccessResponse(c, code, gin.H{"id": id})
}

func (pc *ProductController) GetList(c *gin.Context) {
	products, code := pc.productService.ListProducts()
	response.SuccessResponse(c, code, products)
}

func (pc *ProductController) GetByID(c *gin.Context) {
	product, code := pc.productService.GetProductByID(1)
	response.SuccessResponse(c, code, product)
}

func (pc *ProductController) Update(c *gin.Context) {
	code := pc.productService.UpdateProduct(1, "Coca Chai", "12000")
	response.SuccessResponse(c, code, nil)
}

func (pc *ProductController) Delete(c *gin.Context) {
	code := pc.productService.DeleteProduct(1)
	response.SuccessResponse(c, code, nil)
}
