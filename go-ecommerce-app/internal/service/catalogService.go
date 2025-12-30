package service

import (
	"errors"
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
