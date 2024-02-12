package locations

import (
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"

	"github.com/labstack/echo/v4"
)

type Api struct {
	oapicodegen.ServerInterface
}

func RegisterHandlers(e *echo.Echo) {
	oapicodegen.RegisterHandlers(e, &Api{})
}
