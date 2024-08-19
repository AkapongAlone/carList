package repositories

import (
	"Daveslist/models"
	"Daveslist/src/auth/handlers"
	"Daveslist/src/auth/usecases"

	"gorm.io/gorm"
)

type authRepository struct {
	conn *gorm.DB
}

func NewHandler(conn *gorm.DB) *handlers.Handler {
	usecase := usecases.NewUsecase(&authRepository{conn: conn})
	handler := handlers.NewHandler(usecase)
	return handler
}

func (t *authRepository) GetUserByUserName(username string) (models.User, error) {
	user := models.User{}
	err := t.conn.Where("user_name = ?", username).Find(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
