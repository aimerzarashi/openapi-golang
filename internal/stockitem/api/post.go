package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
	"openapi/internal/stockitem/usecase"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
func Post(c echo.Context) error {
	req := &oapicodegen.PostStockItemJSONBody{}
	c.Bind(&req)

	reqDto := &usecase.CreateRequestDto{
		Name: req.Name,
	}

	resDto, err := usecase.Create(reqDto)
	if err != nil {
		return err
	}

	res := &oapicodegen.Created{
		Id: resDto.Id,
	}

	return c.JSON(http.StatusCreated, res)
}