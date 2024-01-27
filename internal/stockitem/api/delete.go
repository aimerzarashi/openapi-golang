package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
	"openapi/internal/stockitem/usecase"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock item.
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