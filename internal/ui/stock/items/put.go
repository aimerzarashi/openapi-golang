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

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
func Put(c echo.Context) error {
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

	req := &oapicodegen.PutStockItemJSONRequestBody{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqDto := &stockitem.UpdateRequestDto{
		Id:   stockitemId,
		Name: req.Name,
	}

	_, err = stockitem.Update(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := &oapicodegen.OK{}

	return c.JSON(http.StatusOK, res)
}