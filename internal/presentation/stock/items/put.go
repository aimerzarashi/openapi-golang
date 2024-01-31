package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/application/stock/item"
	domain "openapi/internal/domain/stock/item"
	"openapi/internal/infrastructure/database"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
func Put(c echo.Context) error {
	// Pre Process
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Binding
	stockitemId := uuid.MustParse(c.Param("stockitemId"))

	req := &oapicodegen.PutStockItemJSONRequestBody{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validation
	if stockitemId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid stock item id")
	}

	found, err := repository.Find(domain.Id(stockitemId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock item not found")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	reqDto := &item.UpdateRequestDto{
		Id:   stockitemId,
		Name: req.Name,
	}

	resDto, err := item.Update(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Post Process
	res := &oapicodegen.OK{
		Id: resDto.Id,
		Name: resDto.Name,
	}

	return c.JSON(http.StatusOK, res)
}