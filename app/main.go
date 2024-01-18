package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // インスタンスを作成
  e := echo.New()

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルートを設定
  e.GET("/", hello) // ローカル環境の場合、http://localhost:3000/ にGETアクセスされるとhelloハンドラーを実行する

  // サーバーをポート番号3000で起動
  e.Logger.Fatal(e.Start(":3000"))
}

type ResponseData struct {
	Message string `json:"message"`
}

// ハンドラーを定義
func hello(c echo.Context) error {
	data := ResponseData{
		Message: "Hello World",
	}
	return c.JSON(http.StatusOK, data)
}