package main

import (
	"testing"

	"encoding/json"
	"io"
	"net/http"
)

// http://localhost:3000/ にGETでアクセスし、戻り値を検証する
func TestHello(t *testing.T) {
	res, err := http.Get("http://localhost:3000/")
	if err != nil {
			t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
			t.Errorf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	resBodyByte, _ := io.ReadAll(res.Body)
	var data ResponseData = ResponseData{}
	json.Unmarshal(resBodyByte, &data)

	var expect ResponseData = ResponseData{
		Message: "Hello, World!",
	}

	if data != expect {
			t.Errorf("want %s, got %s", "Hello, World!", data.Message)
	}
}