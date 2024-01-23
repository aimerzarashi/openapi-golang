package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"openapi/internal/presentation/hello_api"
)

func GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &hello_api.Hello{
		Message: "Hello, World!",
	})
}