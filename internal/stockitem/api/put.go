package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
func Put(c echo.Context) error {

	response := &oapicodegen.OK{}

	return c.JSON(http.StatusOK, response)
}