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
}
type catelogRepository struct {
	db *gorm.DB
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
