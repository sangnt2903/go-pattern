package config

import (
	"github.com/gin-gonic/gin"
	"go-pattern/controllers"
	"go-pattern/repositories"
	"go-pattern/services"
)

var Container *MasterContainer

type MasterContainer struct {
	UserController *controllers.UserController
}

func InitMasterContainer(r *gin.Engine) {
	// get connection pool
	dbConnection := GetDB()

	// base
	baseRepo := repositories.NewBaseRepository(dbConnection)

	// user
	userRepo := repositories.NewUserRepository(baseRepo)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	controllers.MakeUserHandlerFunc(r, userController)

	Container = &MasterContainer{
		UserController: userController,
	}
}
