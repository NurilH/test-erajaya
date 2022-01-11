package controllers

import (
	"net/http"
	"test_erajaya/lib/databases"
	"test_erajaya/models"

	response "test_erajaya/responses"

	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan user (registrasi) next
func CreateUserControllers(c echo.Context) error {
	new_user := models.Users{}
	c.Bind(&new_user)

	data, err := databases.CreateUser(&new_user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func LoginUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	data, e := databases.LoginUser(&user)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Email or Password Incorrect"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Login Success", data))
}
