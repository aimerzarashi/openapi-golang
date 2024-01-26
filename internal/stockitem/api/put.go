package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
func Put(c echo.Context) error {

	response := &oapicodegen.OK{}

	db, err := database.New()
	if err != nil {
		return err
	}
	defer db.Close()

	return c.JSON(http.StatusOK, response)
}