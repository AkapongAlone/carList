package usecases

import (
	"Daveslist/models"
	"Daveslist/requests"
	"Daveslist/responses"
	"Daveslist/src/messages/domains"

	"github.com/jinzhu/copier"
)

type messageUseCase struct {
	messageRepo domains.Repo
}

func NewUsecase(repo domains.Repo) domains.Usecase {
	return &messageUseCase{messageRepo: repo}
}

func (u *messageUseCase) SendMessage(req requests.MessageRequest, userID uint) error {
	return u.messageRepo.CreateMessage(req, userID)
}

func (u *messageUseCase) ReadMessage(req requests.ReadMessageRequest) ([]responses.MessageResp, error) {
	messageData, err := u.messageRepo.GetMessage(req)
	if err != nil {
		return nil, err
	}

	resp := praseModelMessageToRespMessage(messageData)
	return resp, nil
}

func praseModelMessageToRespMessage(model []models.Message) []responses.MessageResp {
	resp := []responses.MessageResp{}
	copier.Copy(&resp, &model)
	return resp
}
