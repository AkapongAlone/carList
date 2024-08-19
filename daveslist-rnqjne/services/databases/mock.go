package databases

import (
	"Daveslist/constants"
	"Daveslist/models"
)

func MockRole() {
	role := []models.Role{
		{ID: constants.REGISTERED, Name: "registered"}, {ID: constants.MODERATORS, Name: "moderators"}, {ID: constants.SUPERUSER, Name: "superuser"},
	}
	DB.Create(&role)
}

func MockUser() {
	generalUser1 := models.User{UserName: "Akapong", Password: "1234", RoleID: 1}
	generalUser2 := models.User{UserName: "The Gang", Password: "5678", RoleID: 1}
	moderators := models.User{UserName: "moderators", Password: "moderators", RoleID: 3}
	admin := models.User{UserName: "admin", Password: "admin", RoleID: 2}
	userList := []models.User{generalUser1, generalUser2, admin, moderators}
	DB.Create(&userList)
}
