package items

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/application/stock/item"
	domain "openapi/internal/domain/stock/item"
	"openapi/internal/infrastructure/database"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
)

// PostStockItem is a function that handles the HTTP POST request for creating a new stock item.
func Post(c echo.Context) error {
	// Pre Process
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Binding
	req := &oapicodegen.PostStockItemJSONRequestBody{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validation
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	reqDto := &item.CreateRequestDto{
		Name: req.Name,
	}
	resDto, err := item.Create(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Post Process
	res := &oapicodegen.Created{
		Id: resDto.Id,
		Name: resDto.Name,
	}
	return c.JSON(http.StatusCreated, res)
}