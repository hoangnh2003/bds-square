package repo

import (
	"bds-square-backend/internal/database"
	"log"
)

type IProductRepository interface {
	CreateProduct(arg database.CreateProductParams) (int32, error)
	GetProductByID(id int32) (database.Product, error)
	ListProducts() ([]database.Product, error)
	UpdateProduct(arg database.UpdateProductParams) error
	DeleteProduct(id int32) error
}

type productRepository struct{}

func NewProductRepository() IProductRepository {
	return &productRepository{}
}

// CreateProduct implements IProductRepository.
func (pr *productRepository) CreateProduct(arg database.CreateProductParams) (int32, error) {
	result, err := Dao.CreateProduct(Ctx, arg)
	if err != nil {
		log.Println(">>> ERROR: CreateProduct():", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(">>> ERROR: Lấy ID sản phẩm mới thất bại:", err)
		return 0, err
	}
	return int32(id), nil
}

// DeleteProduct implements IProductRepository.
func (pr *productRepository) DeleteProduct(id int32) error {
	err := Dao.DeleteProduct(Ctx, id)
	if err != nil {
		log.Println(">>> ERROR: DeleteProduct():", err)
	}
	return err
}

// GetProductByID implements IProductRepository.
func (pr *productRepository) GetProductByID(id int32) (database.Product, error) {
	product, err := Dao.GetProduct(Ctx, id)
	if err != nil {
		log.Println(">>> ERROR: GetProductByID():", err)
	}
	return product, err
}

// ListProducts implements IProductRepository.
func (r *productRepository) ListProducts() ([]database.Product, error) {
	products, err := Dao.ListProducts(Ctx)
	if err != nil {
		log.Println(">>> ERROR: ListProducts():", err)
	}
	return products, err
}

// UpdateProduct implements IProductRepository.
func (pr *productRepository) UpdateProduct(arg database.UpdateProductParams) error {
	err := Dao.UpdateProduct(Ctx, arg)
	if err != nil {
		log.Println(">>> ERROR: UpdateProduct():", err)
	}
	return err
}
