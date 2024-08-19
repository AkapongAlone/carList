package repositories

import (
	"Daveslist/models"
	"Daveslist/src/categories/handlers"
	"Daveslist/src/categories/usecases"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewHandler(conn *gorm.DB) *handlers.Handler {
	usecase := usecases.NewUsecase(&categoryRepository{conn: conn})
	handler := handlers.NewHandler(usecase)
	return handler
}


func (r *categoryRepository) CreateCategory(data *models.Categories) error {
	return r.conn.Create(data).Error
}


func (r *categoryRepository) DeleteCategory(categoryID uint) error {
	var err error
	tx := r.conn.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()
	err = tx.Where("id = ?", categoryID).Delete(&models.Categories{}).Error
	if err != nil {
		return err
	}
	err = tx.Model(&models.CarList{}).Where("categories_id = ?", categoryID).Update("is_trash", true).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

