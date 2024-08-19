package domains

import (
	"Daveslist/models"
	"Daveslist/requests"
)

type Usecase interface {
	CreateCategory(req requests.RequestCategory, userID uint) error
	DeleteCategory(categoryID uint) error
}

type Repo interface {
	CreateCategory(data *models.Categories) error
	DeleteCategory(categoryID uint) error
}
