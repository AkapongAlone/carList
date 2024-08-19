package usecases

import (
	"Daveslist/helpers"
	"Daveslist/models"
	"Daveslist/requests"
	"Daveslist/src/carList/domains"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type carUseCase struct {
	carRepo domains.Repo
}

func NewUsecase(repo domains.Repo) domains.Usecase {
	return &carUseCase{carRepo: repo}
}

func (u *carUseCase) CreateSellingCar(req requests.RequestSellingCar, c *gin.Context) error {
	var err error
	userID := helpers.GetUserID(c)
	tx := u.carRepo.StartTranSection()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()
	newCarList := models.CarList{
		CreatedBy:    userID,
		CategoriesID: req.CategoryID,
		Title:        req.Title,
		Content:      req.Content,
		IsPublic:     req.IsPublic,
		IsShow:       true,
	}
	err = u.carRepo.CreateCarList(&newCarList, tx)
	if err != nil {
		return err
	}
	listNewImg, err := uploadImageToLocal(req.File, c, newCarList.ID)
	if err != nil {
		return err
	}
	err = u.carRepo.CreateFile(listNewImg, tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func uploadImageToLocal(files []*multipart.FileHeader, c *gin.Context, carListID uint) ([]models.File, error) {
	listNewImg := []models.File{}
	storages := os.Getenv("STORAGES_FOLDER_IMAGE_PATH") + "/"

	if _, err := os.Stat(storages); os.IsNotExist(err) {
		err := os.MkdirAll(storages, 0755)
		if err != nil {
			return nil, err
		}
	}
	for _, file := range files {
		isImg, err := helpers.CheckIsImage(file)
		if err != nil {
			return nil, err
		}
		if isImg {

			URL := filepath.Join(storages, (helpers.GenerateUUIDV4() + file.Filename))
			if err := c.SaveUploadedFile(file, URL); err != nil {
				return nil, err
			}
			newImage := models.File{
				URL:       URL,
				CarListID: carListID,
			}
			listNewImg = append(listNewImg, newImage)
		} else {
			return nil, errors.New("file not image")
		}
	}

	return listNewImg, nil
}

func (u *carUseCase) CreateReply(rep requests.Reply, userID uint) error {
	carList, err := u.carRepo.GetCarListByID(rep.CarListID)
	if err != nil {
		return err
	}

	if checkTimeToReply(carList.CreatedAt) {
		reply := models.Reply{
			Text:      rep.Text,
			CarListID: rep.CarListID,
			CreatedBy: userID,
		}
		err := u.carRepo.CreateReply(&reply)
		if err != nil {
			return err
		}
	} else {
		return errors.New("time expired to reply")
	}
	return nil
}

func (u *carUseCase) EditCarList(req requests.RequestSellingCar, c *gin.Context, carListID uint) error {
	var err error

	tx := u.carRepo.StartTranSection()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()
	err = u.carRepo.DeleteFile(carListID, tx)
	if err != nil {
		return err
	}
	editedCarList, err := u.carRepo.GetCarListByID(carListID)
	if err != nil {
		return err
	}
	editedCarList.CategoriesID = req.CategoryID
	editedCarList.Content = req.Content
	editedCarList.Title = req.Title
	editedCarList.IsPublic = req.IsPublic
	editedCarList.IsShow = req.IsPublic

	err = u.carRepo.UpdateCarList(&editedCarList, tx)
	if err != nil {
		return err
	}
	listNewImg, err := uploadImageToLocal(req.File, c, carListID)
	if err != nil {
		return err
	}

	err = u.carRepo.CreateFile(listNewImg, tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func checkTimeToReply(carListTime time.Time) bool {
	oneYearLater := carListTime.AddDate(1, 0, 0)
	return time.Now().Before(oneYearLater)
}

func (u *carUseCase) GetCarListAndCategories(isRegister bool) ([]models.CategoriesPreload, error) {

	result, err := u.carRepo.GetCarListAndCategories(isRegister)
	if err != nil {
		return nil, err
	}
	baseURL := os.Getenv("BASE_URL")

	for catIndex, category := range result {
		for carIndex, carInfo := range category.CarList {
			for imgIndex, file := range carInfo.Files {
				// Update the file URL
				file.URL = baseURL + "/" + file.URL

				// Assign the updated file back to the result
				result[catIndex].CarList[carIndex].Files[imgIndex] = file
			}
		}
	}

	return result, nil
}

func (u *carUseCase) DeleteList(carListID uint) error {
	return u.carRepo.DeleteList(carListID)
}

func (u *carUseCase) HideCarList(req requests.RequestHideList, carListID uint) error {
	var err error

	editedCarList, err := u.carRepo.GetCarListByID(carListID)
	if err != nil {
		return err
	}

	editedCarList.IsShow = req.IsShow

	err = u.carRepo.UpdateCarList(&editedCarList, nil)
	if err != nil {
		return err
	}

	return nil
}
