package controllers

import (
	"net/http"
	"test_erajaya/lib/databases"
	"test_erajaya/middlewares"
	"test_erajaya/models"
	response "test_erajaya/responses"

	"github.com/labstack/echo/v4"
)

func CreateProductControllers(c echo.Context) error {
	new_product := models.Products{}
	c.Bind(&new_product)

	user_id := middlewares.ExtractTokenId(c)
	new_product.UsersID = uint(user_id)

	data, err := databases.CreateProduct(&new_product)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}
