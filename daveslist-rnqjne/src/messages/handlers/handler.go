package handlers

import (
	"Daveslist/helpers"
	"Daveslist/requests"
	"Daveslist/responses"
	"Daveslist/src/messages/domains"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	messageUseCase domains.Usecase
}

func NewHandler(usecase domains.Usecase) *Handler {
	return &Handler{messageUseCase: usecase}
}

// @summary
// @description
// @tags message
// @id sendMessage
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param SendMessage body requests.MessageRequest true "Login data"
// @response 201 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 422 {object} responses.Validate "Unprocessable Entity"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /message [post]
func (t *Handler) SendMessage(c *gin.Context) {
	var request requests.MessageRequest
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	userID := helpers.GetUserID(c)

	err = t.messageUseCase.SendMessage(request, userID)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)

		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(responses.NoData{}, http.StatusCreated))
}

// @summary
// @description
// @tags message
// @id readMessage
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param from_user_id query string true "from message request"
// @param to_user_id query string true "tp message request"
// @response 200 {object} []responses.MessageResp "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 422 {object} responses.Validate "Unprocessable Entity"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /message [GET]
func (t *Handler) ReadMessage(c *gin.Context) {

	fromID := c.Query("from_user_id")
	toID := c.Query("to_user_id")

	fromIDInt, _ := strconv.Atoi(fromID)

	toIDInt, _ := strconv.Atoi(toID)

	request := requests.ReadMessageRequest{
		FromUserID: uint(fromIDInt),
		ToUserID:   uint(toIDInt),
	}
	userID := helpers.GetUserID(c)
	if userID != request.FromUserID && userID != request.ToUserID {
		errResponse := responses.NewAppErr(http.StatusForbidden, "this user can't read these message")
		c.JSON(http.StatusForbidden, errResponse)
		return
	}
	listMessage, err := t.messageUseCase.ReadMessage(request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)

		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(listMessage, http.StatusCreated))
}
