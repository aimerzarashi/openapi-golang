package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
func Post(c echo.Context) error {

	response := &oapicodegen.Created{
		Id: uuid.New(),
	}

	return c.JSON(http.StatusCreated, response)
}