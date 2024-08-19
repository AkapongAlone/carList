package handlers

import (
	"Daveslist/requests"
	"Daveslist/responses"
	"Daveslist/src/auth/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authUseCase domains.Usecase
}

func NewHandler(usecase domains.Usecase) *Handler {
	return &Handler{authUseCase: usecase}
}

// @summary
// @description
// @tags authentication
// @id login
// @param Login body requests.LoginRequest true "Login data"
// @response 200 {object} responses.TokenResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 422 {object} responses.Validate "Unprocessable Entity"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /auth/login [post]
func (t *Handler) Login(c *gin.Context) {
	// request form
	var request requests.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// call usecase
	accessToken, err := t.authUseCase.Login(request.Username, request.Password)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)

		return
	}

	var loginResponse responses.LoginResponse
	loginResponse.AccessToken = accessToken
	c.JSON(http.StatusOK, responses.SuccessResponse(loginResponse, http.StatusOK))
}
