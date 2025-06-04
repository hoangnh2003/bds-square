package service

import (
	"bds-square-backend/internal/database"
	"bds-square-backend/internal/repo"
	"bds-square-backend/pkg/response"
)

type IProductService interface {
	CreateProduct(name string, price string) (int32, int)
	GetProductByID(id int32) (database.Product, int)
	ListProducts() ([]database.Product, int)
	UpdateProduct(id int32, name string, price string) int
	DeleteProduct(id int32) int
}

type productService struct {
	repo repo.IProductRepository
}

func NewProductService(productRepo repo.IProductRepository) IProductService {
	return &productService{repo: productRepo}
}

// CreateProduct implements IProductService.
func (ps *productService) CreateProduct(name string, price string) (int32, int) {
	id, err := ps.repo.CreateProduct(database.CreateProductParams{
		Name:  name,
		Price: price,
	})
	if err != nil {
		return 0, response.ErrCodeServerError
	}
	return id, response.ErrCodeSuccess
}

// DeleteProduct implements IProductService.
func (ps *productService) DeleteProduct(id int32) int {
	err := ps.repo.DeleteProduct(id)
	if err != nil {
		return response.ErrCodeServerError
	}
	return response.ErrCodeSuccess
}

// GetProductByID implements IProductService.
func (ps *productService) GetProductByID(id int32) (database.Product, int) {
	product, err := ps.repo.GetProductByID(id)
	if err != nil {
		return database.Product{}, response.ErrCodeNotFound
	}
	return product, response.ErrCodeSuccess
}

// ListProducts implements IProductService.
func (ps *productService) ListProducts() ([]database.Product, int) {
	products, err := ps.repo.ListProducts()
	if err != nil {
		return nil, response.ErrCodeServerError
	}
	return products, response.ErrCodeSuccess
}

// UpdateProduct implements IProductService.
func (ps *productService) UpdateProduct(id int32, name string, price string) int {
	err := ps.repo.UpdateProduct(database.UpdateProductParams{
		ID:    id,
		Name:  name,
		Price: price,
	})
	if err != nil {
		return response.ErrCodeServerError
	}
	return response.ErrCodeSuccess
}
