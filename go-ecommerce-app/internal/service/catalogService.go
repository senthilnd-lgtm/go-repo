package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
)

type CatelogService struct {
	Repo   repository.CatelogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatelogService) CreateCategory(input dto.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (s CatelogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {

	existCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("Category does not exist")
	}

	if len(input.Name) > 0 {
		existCat.Name = input.Name
	}

	if input.ParentId > 0 {
		existCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		existCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		existCat.DisplayOrder = input.DisplayOrder
	}

	cat, err := s.Repo.EditCategory(existCat)

	if err != nil {
		return nil, errors.New("Failed to edit category")
	}
	return cat, err
}

func (s CatelogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		return errors.New("Category does not exist to delete")
	}
	return nil
}

func (s CatelogService) GetCategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()

	if err != nil {
		return nil, errors.New("Categogies does not exist")
	}

	return categories, err
}

func (s CatelogService) GetCategory(id int) (*domain.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("Category does not exist")
	}

	return cat, nil
}

func (s CatelogService) CreateProduct(input dto.CreateProductRequest, user domain.User) error {
	err := s.Repo.CreateProduct(&domain.Product{
		Name:        input.Name,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		UserId:      int(user.ID),
		Stock:       uint(input.Stock),
	})

	return err
}

func (s CatelogService) EditProduct(id int, input dto.CreateProductRequest, user domain.User) (*domain.Product, error) {
	exitProduct, err := s.Repo.FindProductById(id)

	if err != nil {
		return nil, errors.New("Product does not exist")
	}

	// verify product owner

	if exitProduct.UserId != int(user.ID) {
		return nil, errors.New("You have not mange this product")
	}

	if len(input.Name) > 0 {
		exitProduct.Name = input.Name
	}

	if len(input.Description) > 0 {
		exitProduct.Description = input.Description
	}

	if input.CategoryId > 0 {
		exitProduct.CategoryId = input.CategoryId
	}

	if input.Price > 0 {
		exitProduct.Price = input.Price
	}

	updateProduct, err := s.Repo.EditProduct(exitProduct)
	return updateProduct, err
}

func (s CatelogService) DeleteProduct(id int, user domain.User) error {
	exitProduct, err := s.Repo.FindProductById(id)

	if err != nil {
		return errors.New("Product does not exist")
	}

	if exitProduct.UserId != int(user.ID) {
		return errors.New("You have not rights to delete this product")
	}

	err = s.Repo.DeleteProduct(exitProduct)

	if err != nil {
		return errors.New("Product cannot delete")
	}
	return nil
}

func (s CatelogService) GetProducts() ([]*domain.Product, error) {
	products, err := s.Repo.FindProducts()
	if err != nil {
		return nil, errors.New("Products does not exist")
	}
	return products, err
}

func (s CatelogService) GetProductById(id int) (*domain.Product, error) {
	product, err := s.Repo.FindProductById(id)
	fmt.Println("SVC product", product)
	if err != nil {
		return nil, errors.New("Product does not exist")
	}
	return product, nil

}

func (s CatelogService) GetSellerProducts(id int) ([]*domain.Product, error) {
	products, err := s.Repo.FindSellerProducts(id)
	if err != nil {
		return nil, errors.New("seller product does not exist")
	}
	return products, err
}

func (s CatelogService) UpdateProductStock(e domain.Product) (*domain.Product, error) {
	product, err := s.Repo.FindProductById(int(e.ID))
	if err != nil {
		return nil, errors.New("Product does not exist")
	}

	// verify product owner

	if product.UserId != e.UserId {
		return nil, errors.New("You dont have manage rights of this product")
	}
	product.Stock = e.Stock
	editProduct, err := s.Repo.EditProduct(product)
	if err != nil {
		return nil, err
	}
	return editProduct, nil
}
