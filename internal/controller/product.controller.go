package controller

import (
	"bds-square-backend/internal/request"
	"bds-square-backend/internal/service"
	"bds-square-backend/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService: productService}
}

// GetList godoc
// @Summary      Get list of products
// @Description  Retrieve the list of all available products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   database.Product
// @Router       /products [get]
func (pc *ProductController) GetList(c *gin.Context) {
	products, code := pc.productService.ListProducts()
	response.SuccessResponse(c, code, products)
}

// GetByID godoc
// @Summary      Get product by ID
// @Description  Retrieve a single product using its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  database.Product
// @Router       /products/{id} [get]
func (pc *ProductController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, code := pc.productService.GetProductByID(int32(id))
	response.SuccessResponse(c, code, product)
}

// Create godoc
// @Summary      Create a new product
// @Description  Add a new product with name and price
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      request.ProductCreateRequest  true  "Product to create"
// @Success      201      {object}  map[string]int32  "Returns the ID of the created product"
// @Router       /products [post]
func (pc *ProductController) Create(c *gin.Context) {
	var req request.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, code := pc.productService.CreateProduct(req.Name, req.Price)
	response.SuccessResponse(c, code, gin.H{"id": id})
}

// Update godoc
// @Summary      Update an existing product
// @Description  Update product information by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Product ID"
// @Param        product  body      request.ProductUpdateRequest   true  "Updated product info"
// @Success      200      {object}  nil
// @Router       /products/{id} [put]
func (pc *ProductController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req request.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	code := pc.productService.UpdateProduct(int32(id), req.Name, req.Price)
	response.SuccessResponse(c, code, nil)
}

// Delete godoc
// @Summary      Delete a product
// @Description  Delete a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {string}  string  "Deleted successfully"
// @Router       /products/{id} [delete]
func (pc *ProductController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	code := pc.productService.DeleteProduct(int32(id))
	response.SuccessResponse(c, code, "Deleted successfully")
}
