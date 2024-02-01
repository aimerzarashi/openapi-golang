package stock

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"openapi/internal/presentation/stock/items"

	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/labstack/echo/v4"
)

type Api struct {}

func New() oapicodegen.ServerInterface {
	return &Api{}
}

func RegisterHandlers(e *echo.Echo, si oapicodegen.ServerInterface) {
	oapicodegen.RegisterHandlers(e, si)
}

func (a *Api) PostStockItem(ctx echo.Context) error {
	return items.PostStockItem(ctx)
}

func (a *Api) PutStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error {
	return items.PutStockItem(ctx, stockitemId)
}

func (a *Api) DeleteStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error {
	return items.DeleteStockItem(ctx, stockitemId)
}