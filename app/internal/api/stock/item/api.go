package stock_item_api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Server struct{
	ServerInterface
}

func (Server) PostStockItem(ctx echo.Context) error {
	request := &PostStockItemJSONBody{}
	ctx.Bind(&request)
	fmt.Println(request)
	
	response := new(CreatedResponse)
	response.Id = uuid.New()
	return ctx.JSON(http.StatusOK, response)
}