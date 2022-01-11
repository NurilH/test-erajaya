package routes

import (
	"test_erajaya/constants"
	"test_erajaya/controllers"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// route users tanpa JWT
	e.POST("/signup", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUserControllers)

	e.GET("/products/sort/newest", controllers.GetProductSortByLastCreateControllers)
	e.GET("/products/sort/asc/name", controllers.GetProductByNameSortAscControllers)
	e.GET("/products/sort/asc/price", controllers.GetProductByPriceSortAscControllers)
	e.GET("/products/sort/desc/name", controllers.GetProductByNameSortDescControllers)
	e.GET("/products/sort/desc/price", controllers.GetProductByPriceSortDescControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	j.POST("/products/create", controllers.CreateProductControllers)

	return e
}
