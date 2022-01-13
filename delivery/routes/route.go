package routes

import (
	"Alterra/batch5/docker-be5/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController) {
	e.GET("/users", uc.Get())
	e.POST("/users/register", uc.CreateUserController())
}
