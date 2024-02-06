package locations

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infrastructure/database"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func DeleteStockLocation(ctx echo.Context, stockLocationId openapi_types.UUID) error {
	// Preprocess
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Preprocess Validation
	if stockLocationId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid stock location id")
	}

	found, err := repository.Find(domain.Id(stockLocationId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock location not found")
	}

	// Main Process
	reqDto := &location.DeleteRequestDto{
		Id: stockLocationId,
	}
	if err := location.Delete(reqDto, repository); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}