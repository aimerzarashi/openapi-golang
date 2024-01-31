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

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func Delete(c echo.Context) error {
	// Pre Process
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Binding
	stockitemId := uuid.MustParse(c.Param("stockitemId"))

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
	
	// Main Process
	reqDto := &item.DeleteRequestDto{
		Id:   stockitemId,
	}
	_, err = item.Delete(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Post Process
	res := &oapicodegen.OK{}

	return c.JSON(http.StatusOK, res)
}