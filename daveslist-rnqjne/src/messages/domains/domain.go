package domains

import (
	"Daveslist/models"
	"Daveslist/requests"
	"Daveslist/responses"
)

type Usecase interface {
	SendMessage(req requests.MessageRequest,userID uint) error
	ReadMessage(req requests.ReadMessageRequest) ([]responses.MessageResp, error)
}

type Repo interface {
	CreateMessage(req requests.MessageRequest, userID uint) error
	GetMessage(req requests.ReadMessageRequest) ([]models.Message, error)
}
