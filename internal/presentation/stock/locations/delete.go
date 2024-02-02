package locations

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infrastructure/database"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func DeleteStockLocation(ctx echo.Context, stockitemId openapi_types.UUID) error {
	// Preprocess
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}