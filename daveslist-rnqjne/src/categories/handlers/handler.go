package handlers

import (
	"Daveslist/helpers"
	"Daveslist/requests"
	"Daveslist/responses"
	"Daveslist/src/categories/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	categoryUseCase domains.Usecase
}

func NewHandler(usecase domains.Usecase) *Handler {
	return &Handler{categoryUseCase: usecase}
}


// @summary
// @description
// @tags category
// @id create-category
// @param category_data body requests.RequestCategory true "category data data"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @response 201 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/category [post]
func (t *Handler) CreateCategory(c *gin.Context) {
	
	var request requests.RequestCategory
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	userID := helpers.GetUserID(c)
	
	err = t.categoryUseCase.CreateCategory(request, userID)
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(responses.NoData{}, http.StatusCreated))
}


// @summary
// @description
// @tags category
// @id delete-category
// @param id path string true "category id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @response 204 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/category/{id} [delete]
func (t *Handler) DeleteCategory(c *gin.Context) {

	categoryID, err := helpers.GetIDParam(c)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	// call usecase
	err = t.categoryUseCase.DeleteCategory(uint(categoryID))
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusNoContent, responses.SuccessResponse(responses.NoData{}, http.StatusNoContent))
}

