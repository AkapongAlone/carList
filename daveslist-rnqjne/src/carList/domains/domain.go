package domains

import (
	"Daveslist/models"
	"Daveslist/requests"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Usecase interface {
	GetCarListAndCategories(IsRegister bool) ([]models.CategoriesPreload, error)
	CreateSellingCar(req requests.RequestSellingCar, c *gin.Context) error
	EditCarList(req requests.RequestSellingCar, c *gin.Context, carListID uint) error
	HideCarList(req requests.RequestHideList, carListID uint) error
	CreateReply(rep requests.Reply, userID uint) error

	DeleteList(categoryID uint) error
}

type Repo interface {
	CreateCarList(data *models.CarList, tx *gorm.DB) error
	CreateFile(data []models.File, tx *gorm.DB) error
	GetCarListByID(carListID uint) (models.CarList, error)
	CreateReply(data *models.Reply) error
	DeleteFile(CarListID uint, tx *gorm.DB) error
	UpdateCarList(editedCarList *models.CarList, tx *gorm.DB) error
	StartTranSection() *gorm.DB
	GetCarListAndCategories(isRegister bool) ([]models.CategoriesPreload, error)
	GetTrashBin() ([]models.CarList, error)
	DeleteList(carListID uint) error
}
