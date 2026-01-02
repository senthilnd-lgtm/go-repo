package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"

	"gorm.io/gorm"
)

type CatelogRepository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryById(id int) (*domain.Category, error)
	EditCategory(e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error

	CreateProduct(e *domain.Product) error
	FindProducts() ([]*domain.Product, error)
	FindProductById(id int) (*domain.Product, error)
	FindSellerProducts(id int) ([]*domain.Product, error)
	EditProduct(e *domain.Product) (*domain.Product, error)
	DeleteProduct(e *domain.Product) error
}
type catelogRepository struct {
	db *gorm.DB
}

// CreateProduct implements CatelogRepository.
func (c *catelogRepository) CreateProduct(e *domain.Product) error {
	err := c.db.Model(&domain.Product{}).Create(e).Error
	if err != nil {
		return errors.New("Cannot create product")
	}
	return nil
}

// DeleteProduct implements CatelogRepository.
func (c *catelogRepository) DeleteProduct(e *domain.Product) error {
	err := c.db.Delete(&domain.Product{}, e.ID).Error

	if err != nil {
		return errors.New("Product cannot delete")
	}
	return nil
}

// EditProduct implements CatelogRepository.
func (c *catelogRepository) EditProduct(e *domain.Product) (*domain.Product, error) {
	err := c.db.Save(&e).Error
	if err != nil {
		return nil, errors.New("Failed to update product")
	}
	return e, nil
}

// FindProductById implements CatelogRepository.
func (c *catelogRepository) FindProductById(id int) (*domain.Product, error) {
	var product *domain.Product
	err := c.db.First(&product, id).Error
	if err != nil {
		return nil, errors.New("Product does not exist")
	}
	return product, nil
}

// FindProducts implements CatelogRepository.
func (c *catelogRepository) FindProducts() ([]*domain.Product, error) {
	var products []*domain.Product

	err := c.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindSellerProducts implements CatelogRepository.
func (c *catelogRepository) FindSellerProducts(id int) ([]*domain.Product, error) {
	var products []*domain.Product
	err := c.db.Where("user_id=?").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// CreateCategory implements CatelogRepository.
func (c *catelogRepository) CreateCategory(e *domain.Category) error {
	err := c.db.Create(&e).Error

	if err != nil {
		return errors.New("crate category failed")
	}
	return nil
}

// DeleteCategory implements CatelogRepository.
func (c *catelogRepository) DeleteCategory(id int) error {
	err := c.db.Delete(&domain.Category{}, id).Error

	if err != nil {
		return err
	}

	return nil

}

// EditCategory implements CatelogRepository.
func (c *catelogRepository) EditCategory(e *domain.Category) (*domain.Category, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		return nil, errors.New("Fail to update category")
	}
	return e, nil
}

// FindCategories implements CatelogRepository.
func (c *catelogRepository) FindCategories() ([]*domain.Category, error) {

	var categories []*domain.Category
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// FindCategoryById implements CatelogRepository.
func (c *catelogRepository) FindCategoryById(id int) (*domain.Category, error) {
	var category *domain.Category
	err := c.db.First(&category, id).Error

	if err != nil {
		return nil, errors.New("Category does not exist")
	}
	return category, nil
}

func NewCatelogRepository(db *gorm.DB) CatelogRepository {
	return &catelogRepository{
		db: db,
	}
}
