package domains

import "Daveslist/models"

type Usecase interface {
	Login(username, password string) (string, error)
}

type Repo interface {
	GetUserByUserName(username string) (models.User, error)
}
