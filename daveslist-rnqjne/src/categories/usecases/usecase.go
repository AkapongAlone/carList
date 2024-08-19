package usecases

import (
	"Daveslist/models"
	"Daveslist/requests"
	"Daveslist/src/categories/domains"
)

type categoryUseCase struct {
	carRepo domains.Repo
}

func NewUsecase(repo domains.Repo) domains.Usecase {
	return &categoryUseCase{carRepo: repo}
}

func (u *categoryUseCase) CreateCategory(req requests.RequestCategory, userID uint) error {
	newCategory := models.Categories{
		CreatedBy:    userID,
		CategoryName: req.CategoryName,
		IsPublic:     req.IsPublic,
	}
	err := u.carRepo.CreateCategory(&newCategory)
	if err != nil {
		return err
	}
	return nil
}

func (u *categoryUseCase) DeleteCategory(categoryID uint) error {
	return u.carRepo.DeleteCategory(categoryID)
}

