package middlewares

import (
	"Daveslist/constants"
	"Daveslist/helpers"
	"Daveslist/models"
	"Daveslist/responses"
	"Daveslist/services/databases"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/samber/lo"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		errUnAuth := responses.AppErr{
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Message:    "unauthorized to access the resource",
		}
		// errResponse := responses.FailRespone(&errUnAuth)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errUnAuth)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := validateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errUnAuth)
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errUnAuth)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errUnAuth)
			return
		}

		c.Set("userID", payload["user_id"])
		c.Set("roleID", payload["role_id"])
		// userPermission := []string{}
		// if payload["access_control"] != nil {
		// 	userPermission = helpers.GetUserPermission(payload["access_control"].([]interface{}))
		// }
		// if payload["user_id"] != nil && payload["department_id"] != nil {
		// 	c.Set("userInfo", helpers.UserInfo{
		// 		UserId:         int(payload["user_id"].(float64)),
		// 		UserDept:       int(payload["department_id"].(float64)),
		// 		UserPermission: userPermission,
		// 	})
		// }

		c.Next()
	}
}

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
}

func HasPermissionAccess(allowRoleID ...uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := helpers.GetRoleID(c)
		if !lo.Contains(allowRoleID, roleID) {
			errPermission := responses.AppErr{
				Status:     false,
				StatusCode: http.StatusForbidden,
				Message:    "permission denied",
			}
			c.AbortWithStatusJSON(http.StatusForbidden, errPermission)
			return
		}
		c.Next()
	}
}

func HasPermissionManageList() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := helpers.GetRoleID(c)
		if roleID == constants.REGISTERED {
			userID := helpers.GetUserID(c)
			carListID, err := helpers.GetIDParam(c)
			if err != nil {
				errResponse := responses.NewAppErr(http.StatusBadRequest, err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
				return
			}
			carListInfo := models.CarList{}
			err = databases.DB.Where("id = ? and created_by = ?", carListID, userID).First(&carListInfo).Error
			if err != nil {
				errPermission := responses.AppErr{
					Status:     false,
					StatusCode: http.StatusForbidden,
					Message:    "permission denied",
				}
				c.AbortWithStatusJSON(http.StatusForbidden, errPermission)
				return
			}
		} else if roleID != constants.SUPERUSER {
			errPermission := responses.AppErr{
				Status:     false,
				StatusCode: http.StatusForbidden,
				Message:    "permission denied",
			}
			c.AbortWithStatusJSON(http.StatusForbidden, errPermission)
			return
		}

		c.Next()
	}
}
