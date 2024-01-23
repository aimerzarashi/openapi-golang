package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"openapi/internal/usecase/hello"
	"openapi/internal/usecase/stock_item"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbDriver := "postgres"
	dsn := "host=openapi-db port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, openErr := sql.Open(dbDriver, dsn)
	if openErr != nil {
		e.Logger.Fatal(openErr)
	}
	defer db.Close()

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	e.GET("/", hello.GetHello)
	e.POST("/stock/items", func(ctx echo.Context) error {
		return stock_item.PostStockItem(ctx, db)
	})

	// サーバーをポート番号3000で起動
	e.Logger.Fatal(e.Start(":3000"))
}
