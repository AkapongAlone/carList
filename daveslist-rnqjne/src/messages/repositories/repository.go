package repositories

import (
	"Daveslist/models"
	"Daveslist/requests"
	"Daveslist/src/messages/handlers"
	"Daveslist/src/messages/usecases"

	"gorm.io/gorm"
)

type messageRepository struct {
	conn *gorm.DB
}

func NewHandler(conn *gorm.DB) *handlers.Handler {
	usecase := usecases.NewUsecase(&messageRepository{conn: conn})
	handler := handlers.NewHandler(usecase)
	return handler
}

func (r *messageRepository) CreateMessage(req requests.MessageRequest, userID uint) error {
	newMessage := models.Message{
		Message:    req.Message,
		FromUserID: userID,
		ToUserID:   req.ToUserID,
	}
	return r.conn.Create(&newMessage).Error
}

func (r *messageRepository) GetMessage(req requests.ReadMessageRequest) ([]models.Message, error) {
	listMessages := []models.Message{}
	err := r.conn.Where("from_user_id = ? AND to_user_id = ?", req.FromUserID, req.ToUserID).Find(&listMessages).Order("created_at ASC").Error
	if err != nil {
		return nil, err
	}
	return listMessages, nil
}
