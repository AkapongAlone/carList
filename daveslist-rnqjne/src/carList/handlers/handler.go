package handlers

import (
	"Daveslist/helpers"
	"Daveslist/requests"
	"Daveslist/responses"
	"Daveslist/src/carList/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	carUseCase domains.Usecase
}

func NewHandler(usecase domains.Usecase) *Handler {
	return &Handler{carUseCase: usecase}
}

// @summary Create a new car listing
// @description Create a new car listing with associated files
// @tags car_list
// @id create_car_list
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param category_id formData uint true "Category ID"
// @Param title formData string true "Title"
// @Param content formData string true "Content"
// @Param file formData []file true "Uploaded files"
// @Param is_public formData bool true "Is the listing public?"
// @response 201 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/car_list [post]
func (t *Handler) CreateSellingCar(c *gin.Context) {
	// request form
	var request requests.RequestSellingCar
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	if len(request.File) > 10 {
		errResponse := responses.NewAppErr(http.StatusBadRequest, "over limit image file")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	// call usecase
	err = t.carUseCase.CreateSellingCar(request, c)
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(responses.NoData{}, http.StatusCreated))
}

// @summary
// @description
// @tags car_list
// @id edit_car_list
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param selling body requests.RequestSellingCar true "selling car data"
// @param id path string true "selling car id"
// @response 202 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/car_list/{id} [put]
func (t *Handler) EditSellingCar(c *gin.Context) {
	// request form
	var request requests.RequestSellingCar
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	if len(request.File) > 10 {
		errResponse := responses.NewAppErr(http.StatusBadRequest, "over limit image file")
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	carListID, err := helpers.GetIDParam(c)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = t.carUseCase.EditCarList(request, c, uint(carListID))
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusAccepted, responses.SuccessResponse(responses.NoData{}, http.StatusAccepted))
}

// @summary
// @description
// @tags car_list
// @id delete_car_list
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param id path string true "selling car id"
// @response 204 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/car_list/{id} [delete]
func (t *Handler) DeleteSellingCar(c *gin.Context) {

	carListID, err := helpers.GetIDParam(c)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = t.carUseCase.DeleteList(uint(carListID))
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusNoContent, responses.SuccessResponse(responses.NoData{}, http.StatusNoContent))
}

// @summary
// @description
// @tags car_list
// @id get_sell_car
// @Param Authorization header string false "Insert your access token" default(Bearer <Add access token here>)
// @response 200 {object} responses.Success{data=[]responses.CarCategoriesResp} "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/car_list [get]
func (t *Handler) GetSellingCar(c *gin.Context) {

	isRegister := c.GetBool("userID")
	sellingList, err := t.carUseCase.GetCarListAndCategories(isRegister)
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(sellingList, http.StatusCreated))
}



// @summary
// @description
// @tags car_list
// @id hide_car_list
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param hide body requests.RequestHideList true "want to hide or not"
// @param id path string true "car_list id"
// @response 202 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/car_list/{id}/hide [put]
func (t *Handler) HideCarList(c *gin.Context) {
	var request requests.RequestHideList
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	carListID, err := helpers.GetIDParam(c)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = t.carUseCase.HideCarList(request, uint(carListID))
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(responses.NoData{}, http.StatusCreated))
}


// @summary
// @description
// @tags reply
// @id reply_sell_car
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param reply body requests.Reply true "reply data"
// @param id path string true "selling car id"
// @response 201 {object} responses.NoDataResponse "OK"
// @response 400 {object} responses.BadRequestErrorResponse "Bad Request"
// @response 401 {object} responses.UnauthorizedErrorResponse "Unauthorized"
// @response 500 {object} responses.InternalServerErrorResponse "Internal Server Error"
// @router /sell_car/car_list/{id} [post]
func (t *Handler) ReplySellingCar(c *gin.Context) {

	var request requests.Reply
	err := c.ShouldBind(&request)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	carListID, err := helpers.GetIDParam(c)
	if err != nil {
		errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	request.CarListID = carListID
	userID := helpers.GetUserID(c)
	err = t.carUseCase.CreateReply(request, userID)
	if err != nil {
		errResponse := responses.NewAppErr(500, err.Error())
		c.JSON(500, errResponse)

		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse(responses.NoData{}, http.StatusCreated))
}
