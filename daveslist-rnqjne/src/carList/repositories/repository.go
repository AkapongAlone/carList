package repositories

import (
	"Daveslist/models"
	"Daveslist/src/carList/handlers"
	"Daveslist/src/carList/usecases"

	"gorm.io/gorm"
)

type carRepository struct {
	conn *gorm.DB
}

func NewHandler(conn *gorm.DB) *handlers.Handler {
	usecase := usecases.NewUsecase(&carRepository{conn: conn})
	handler := handlers.NewHandler(usecase)
	return handler
}

func (r *carRepository) CreateCarList(data *models.CarList, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(data).Error
	}
	return r.conn.Create(data).Error
}

func (r *carRepository) CreateFile(data []models.File, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(data).Error
	}
	return r.conn.Create(data).Error
}

func (r *carRepository) CreateCategory(data *models.Categories) error {
	return r.conn.Create(data).Error
}

func (r *carRepository) CreateReply(data *models.Reply) error {
	return r.conn.Create(data).Error
}
func (r *carRepository) GetCarListByID(carListID uint) (models.CarList, error) {
	carList := models.CarList{}
	err := r.conn.Where("id = ?",carListID).First(&carList).Error
	if err != nil {
		return models.CarList{}, err
	}
	return carList, nil
}

func (r *carRepository) DeleteFile(CarListID uint, tx *gorm.DB) error {
	if tx != nil {
		return tx.Where("car_list_id = ?", CarListID).Delete(&models.File{}).Error
	}
	return r.conn.Where("car_list_id = ?", CarListID).Delete(&models.File{}).Error
}

func (r *carRepository) UpdateCarList(editedCarList *models.CarList, tx *gorm.DB) error {
	if tx != nil {
		return tx.Save(editedCarList).Error
	}
	return r.conn.Save(editedCarList).Error
}

func (r *carRepository) StartTranSection() *gorm.DB {
	return r.conn.Begin()
}

func (r *carRepository) GetCarListAndCategories(isRegister bool) ([]models.CategoriesPreload, error) {
	var err error
	result := []models.CategoriesPreload{}
	if isRegister {
		err = r.conn.Preload("CarList", "is_show = ? AND is_trash = ?", true, false).Preload("CarList.Replies").Preload("CarList.Files").Find(&result).Error
	} else {
		err = r.conn.Preload("CarList", "is_public = ? AND is_show = ? AND is_trash = ?", true, true, false).Preload("CarList.Replies").Preload("CarList.Files").Where("is_public = ?", true).Find(&result).Error
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *carRepository) GetTrashBin() ([]models.CarList, error) {
	result := []models.CarList{}
	err := r.conn.Where("is_trash = ?", true).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *carRepository) DeleteList(carListID uint) error {
	return r.conn.Where("id = ?", carListID).Delete(&models.CarList{}).Error
}
