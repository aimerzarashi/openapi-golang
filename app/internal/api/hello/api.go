package hello_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct{
	ServerInterface
}

func (Server) GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &Hello{
		Message: "Hello, World!",
	})
}