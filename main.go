package main

import (
	"Alterra/batch5/docker-be5/configs"
	// "Project/playground/be5/rest/layered/configs"
	"Alterra/batch5/docker-be5/delivery/controllers/user"
	"Alterra/batch5/docker-be5/delivery/routes"
	_userRepo "Alterra/batch5/docker-be5/repository/user"
	"Alterra/batch5/docker-be5/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := _userRepo.New(db)
	userController := user.New(*userRepo)

	e := echo.New()

	routes.RegisterPath(e, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
