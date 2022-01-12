package controllers

import (
	"net/http"
	"regexp"
	"strconv"
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

	if !regexp.MustCompile("^[1-9]?[0-9].*$").MatchString(strconv.Itoa(new_product.Quantity)) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid quantity"))
	}
	if !regexp.MustCompile("^[1-9]?[0-9]{3}.*$").MatchString(strconv.Itoa(new_product.Price)) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Price or price must be >=1000"))
	}
	if !regexp.MustCompile("^[0-9A-Za-z].*$").MatchString(new_product.Name) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Name"))
	}
	if !regexp.MustCompile("^[0-9A-Za-z].*$").MatchString(new_product.Description) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Description"))
	}

	data, err := databases.CreateProduct(&new_product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))

}

func GetProductBySortLastCreateControllers(c echo.Context) error {
	product, err := databases.GetProductBySortLastCreate()
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}

func GetProductByNameSortAscControllers(c echo.Context) error {
	product, err := databases.GetProductByNameSortAsc()
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}

func GetProductByNameSortDescControllers(c echo.Context) error {
	product, err := databases.GetProductByNameSortDesc()
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}

func GetProductByPriceSortAscControllers(c echo.Context) error {
	product, err := databases.GetProductByPriceSortAsc()
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}

func GetProductByPriceSortDescControllers(c echo.Context) error {
	product, err := databases.GetProductByPriceSortDesc()
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", product))
}
