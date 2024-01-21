package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	HelloApi "openapi/internal/api/hello"
	StockItemApi "openapi/internal/api/stock/item"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	helloApi := HelloApi.Server{}
	HelloApi.RegisterHandlers(e, helloApi)
	stockItemApi := StockItemApi.Server{}
	StockItemApi.RegisterHandlers(e, stockItemApi)

	// サーバーをポート番号3000で起動
	e.Logger.Fatal(e.Start(":3000"))
}
