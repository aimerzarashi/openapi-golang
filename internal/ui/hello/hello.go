package hello

import (
	oapicodegen "openapi/internal/infra/oapicodegen/hello"

	"github.com/labstack/echo/v4"
)

type Api struct {}

func RegisterHandlers(e *echo.Echo) {
	oapicodegen.RegisterHandlers(e, &Api{})
}

func (a *Api) GetHello(ctx echo.Context) error {
	return GetHello(ctx)
}