package locations

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

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

	// Postprocess
	res := &oapicodegen.Created{Id: uuid.New()}

	// Postcondition Validation
	if err := ctx.Validate(res); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}