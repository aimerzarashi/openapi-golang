package locations

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/infrastructure/database"

	openapi_types "github.com/oapi-codegen/runtime/types"

	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock location.
func PutStockLocation(ctx echo.Context, stockLocationId openapi_types.UUID) error {
	// Preprocess
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()


	// Binding
	req := &oapicodegen.PutStockLocationJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Preprocess Validation
	if stockLocationId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid stock location id")
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}