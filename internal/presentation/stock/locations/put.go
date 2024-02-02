package locations

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/infrastructure/database"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
func PutStockLocation(ctx echo.Context, stockitemId openapi_types.UUID) error {
	// Preprocess
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}