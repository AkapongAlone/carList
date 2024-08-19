package usecases

import (
	"Daveslist/models"
	"Daveslist/src/auth/domains"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type authUseCase struct {
	authRepo domains.Repo
}

func NewUsecase(repo domains.Repo) domains.Usecase {
	return &authUseCase{authRepo: repo}
}

func (t *authUseCase) Login(username, password string) (string, error) {
	user, err := t.authRepo.GetUserByUserName(username)
	if err != nil {

		return "", err
	}

	if user.UserName != username || user.Password != password || user.ID == 0 {
		return "", errors.New("user or password invalid")
	}

	accessToken, err := createToken(user)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func createToken(user models.User) (string, error) {
	claims := jwt.MapClaims{}

	claims["user_id"] = user.ID
	claims["role_id"] = user.RoleID
	claims["exp"] = time.Now().Add(time.Minute * 3600).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, err
}
