package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
	"openapi/internal/stockitem/domain"
	"openapi/internal/stockitem/repository"
	"openapi/internal/stockitem/usecase"
)

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func Delete(c echo.Context) error {
	stockitemId := uuid.MustParse(c.Param("stockitemId"))
	if stockitemId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid stock item id")
	}

	db, err := database.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	found, err := repository.Find(db, domain.StockItemId(stockitemId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock item not found")
	}
	
	reqDto := &usecase.DeleteRequestDto{
		Id:   stockitemId,
	}

	_, err = usecase.Delete(reqDto, db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := &oapicodegen.OK{}

	return c.JSON(http.StatusOK, res)
}