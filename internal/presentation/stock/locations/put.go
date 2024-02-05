package locations

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
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
	repository := &domain.Repository{Db: db}

	// Binding
	req := &oapicodegen.PutStockLocationJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

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

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	reqDto := &location.UpdateRequestDto{
		Id:   stockLocationId,
		Name: req.Name,
	}
	err = location.Update(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}