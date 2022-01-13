package user

import (
	// "Project/playground/be5/rest/layered/repository/user"
	"net/http"

	"Alterra/batch5/docker-be5/entities"
	"Alterra/batch5/docker-be5/repository/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo user.UserRepository
}

func New(user user.UserRepository) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.Repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Something wrong")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    users,
		})
	}
}
func (uc UserController) CreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := entities.User{}

		if err := c.Bind(&requestFormat); err != nil {

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "There is something wrong with the input",
			})
		}
		newUser := entities.User{
			Nama: requestFormat.Nama,
			HP:   requestFormat.HP,
		}
		res, err := uc.Repo.CreateUserController(newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "There is something wrong",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"users":   res,
		})
	}

}
