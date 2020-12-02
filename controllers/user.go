package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pattern/models"
	"go-pattern/resources"
	"go-pattern/services"
	"net/http"
)

func (uc *UserController) FetchUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := uc.userService.FetchUsers(&users); err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, models.PublicUsers(users))
	}
}

func (uc *UserController) StoreUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request resources.UserInsertRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		var user models.User
		//if err := mapstructure.Decode(request, &user); err != nil {
		//	c.JSON(http.StatusOK, err.Error())
		//	return
		//}
		user = user.Init(request.Email, request.FirstName, request.LastName)
		if err := uc.userService.StoreUser(&user); err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, user.Public())
	}
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{us}
}

type UserController struct {
	userService *services.UserService
}

func MakeUserHandlerFunc(r *gin.Engine, controller *UserController) {
	users := r.Group("/users")
	{
		users.GET("/", controller.FetchUsers())
		users.POST("/", controller.StoreUser())
	}
}
