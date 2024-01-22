package stock_item_api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"openapi/repository/models"
)

type Server struct{
	ServerInterface
}

func (Server) PostStockItem(ctx echo.Context) error {
	request := &PostStockItemJSONBody{}
	ctx.Bind(&request)

	stockItemId := uuid.New()

	stockItem := models.StockItem{}
	stockItem.ID = stockItemId.String()
	stockItem.Name = request.Name
	var stockItemRepository = StockItemRepository{}
	storeErr := stockItemRepository.Store(stockItem)
	if storeErr != nil {
		return ctx.JSON(http.StatusInternalServerError, storeErr.Error())
	}

	response := new(CreatedResponse)
	response.Id = stockItemId
	return ctx.JSON(http.StatusOK, response)
}

type StockItemRepository struct {}

func (s StockItemRepository) Store(stockItem models.StockItem) error {
	dbDriver := "postgres"
	dsn := "host=127.0.0.1 port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, openErr := sql.Open(dbDriver, dsn)
	if openErr != nil {
		return openErr
	}
	defer db.Close()
	
	execErr := stockItem.Insert(context.Background(), db, boil.Infer())
	if execErr != nil {
		return execErr
	}

	return nil
}