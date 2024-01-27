package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
	"openapi/internal/stockitem/usecase"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
func Post(c echo.Context) error {
	req := &oapicodegen.PostStockItemJSONRequestBody{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer db.Close()

	reqDto := &usecase.CreateRequestDto{
		Name: req.Name,
	}

	resDto, err := usecase.Create(reqDto, db)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := &oapicodegen.Created{
		Id: resDto.Id,
	}
	return c.JSON(http.StatusCreated, res)
}