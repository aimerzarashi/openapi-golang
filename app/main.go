package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	oapi "openapi/codegen"
)

type HelloApiController struct{}

// GetHello implements generated.ServerInterface.
func (HelloApiController) GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &oapi.Hello{
		Message: "Hello, World!",
	})
}

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	helloApi := HelloApiController{}
	oapi.RegisterHandlers(e, helloApi)

	// サーバーをポート番号3000で起動
	e.Logger.Fatal(e.Start(":3000"))
}
