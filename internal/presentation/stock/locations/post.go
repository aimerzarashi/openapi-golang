package locations

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infrastructure/database"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
)

// PostStockLocation is a function that handles the HTTP POST request for creating a new stock item.
func PostStockLocation(ctx echo.Context) error {
	// Preprocess
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Binding
	req := &oapicodegen.PostStockLocationJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Precondition Validation
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	reqDto := &location.CreateRequestDto{
		Name: req.Name,
	}
	resDto, err := location.Create(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postprocess
	res := &oapicodegen.Created{Id: resDto.Id}

	// Postcondition Validation
	if err := ctx.Validate(res); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}