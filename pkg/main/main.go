package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	hello "openapi/internal/presentation/hello"
	stockitem "openapi/internal/presentation/stock/items"
)


type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    // Optionally, you could return the error to give each route more control over the status code
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
  e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", hello.Get)
	e.POST("/stock/items", stockitem.Post)
	e.PUT("/stock/items/:stockitemId", stockitem.Put)
	e.DELETE("/stock/items/:stockitemId", stockitem.Delete)
	e.Logger.Fatal(e.Start(":1323"))
}