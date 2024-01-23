package stock_item

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/domain"
	"openapi/internal/presentation/stock_item_api"
	"openapi/internal/repository/stock_item"
)

func PostStockItem(ctx echo.Context, db *sql.DB) error {
	request := stock_item_api.PostStockItemJSONRequestBody{}
	ctx.Bind(&request)

	stockItem := &domain.StockItem{
		Id: uuid.New(),
		Name: request.Name,
	}

	stockItemRepository := &stock_item.Repository{}
	err := stockItemRepository.Store(db, stockItem)
	if err != nil {
		return err
	}

	 reponse := stock_item_api.CreatedResponse{
		Id: stockItem.Id,
	 }

	return ctx.JSON(http.StatusOK, reponse)
}