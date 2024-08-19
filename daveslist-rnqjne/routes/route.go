package routes

import (
	"os"
	"strings"

	"Daveslist/constants"
	_ "Daveslist/docs"
	"Daveslist/middlewares"
	"Daveslist/services/databases"
	authRepo "Daveslist/src/auth/repositories"
	carListRepo "Daveslist/src/carList/repositories"
	categoryRepo "Daveslist/src/categories/repositories"
	messageRepo "Daveslist/src/messages/repositories"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ensureContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // process request first, then modify headers

		switch {
		case strings.HasSuffix(c.Request.URL.Path, ".png"):
			c.Header("Content-Type", "image/png")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		case strings.HasSuffix(c.Request.URL.Path, ".jpg") || strings.HasSuffix(c.Request.URL.Path, ".jpeg"):
			c.Header("Content-Type", "image/jpeg")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	imageFolderPath := os.Getenv("STORAGES_FOLDER_IMAGE_PATH")
	//allowedOrigins := os.Getenv("ALLOW_ORIGIN")

	r.Use(ensureContentType())

	r.StaticFS("/storages/images", gin.Dir(imageFolderPath, true))
	r.Use(func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")

		// Check if the request's origin is in the list of allowed domains

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	message := r.Group("message", middlewares.AuthorizeJWT())
	{
		messageHandler := messageRepo.NewHandler(databases.DB)
		message.POST("", messageHandler.SendMessage)
		message.GET("", messageHandler.ReadMessage)
	}
	auth := r.Group("auth")
	{
		authHandler := authRepo.NewHandler(databases.DB)
		auth.POST("login", authHandler.Login)
	}
	// setup routes
	sellCar := r.Group("sell_car")
	{
		carListHandler := carListRepo.NewHandler(databases.DB)

		carList := sellCar.Group("car_list")
		{
			carList.GET("", carListHandler.GetSellingCar)
			carList.Use(middlewares.AuthorizeJWT())
			carList.POST("", middlewares.HasPermissionAccess(constants.REGISTERED), carListHandler.CreateSellingCar)
			carList.PUT(":id", middlewares.HasPermissionManageList(), carListHandler.EditSellingCar)
			carList.DELETE(":id", middlewares.HasPermissionManageList(), carListHandler.DeleteSellingCar)
			carList.POST(":id", middlewares.HasPermissionAccess(constants.REGISTERED), carListHandler.ReplySellingCar)

		}
		category := sellCar.Group("category", middlewares.AuthorizeJWT())
		{
			categoryHandler := categoryRepo.NewHandler(databases.DB)
			category.POST("", middlewares.HasPermissionAccess(constants.MODERATORS, constants.SUPERUSER), categoryHandler.CreateCategory)
			category.PUT(":id/hide", middlewares.HasPermissionAccess(constants.MODERATORS, constants.SUPERUSER), carListHandler.HideCarList)
			category.DELETE(":id", middlewares.HasPermissionAccess(constants.MODERATORS, constants.SUPERUSER), categoryHandler.DeleteCategory)

		}
	}
	return r
}
