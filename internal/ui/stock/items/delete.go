package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/application/stockitem"
	"openapi/internal/domain/model"
	"openapi/internal/domain/repository"
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
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
	repository := &repository.StockItem{DB: db}

	found, err := repository.Find(model.StockItemId(stockitemId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock item not found")
	}
	
	reqDto := &stockitem.DeleteRequestDto{
		Id:   stockitemId,
	}

	_, err = stockitem.Delete(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := &oapicodegen.OK{}

	return c.JSON(http.StatusOK, res)
}