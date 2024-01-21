package stock_item_api

import (
	"bytes"
	"testing"

	cmp "github.com/google/go-cmp/cmp"
	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

// http://localhost:3000/ にGETでアクセスし、戻り値を検証する
func TestHello(t *testing.T) {
	request := new(PostStockItemJSONBody)
	request.Name = "test"
	requestJson, _ := json.Marshal(request)
	res, err := http.Post("http://localhost:3000/stock/items",
		"application/json",
		bytes.NewBuffer(requestJson))
	if err != nil {
			t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
			t.Errorf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	resBodyByte, _ := io.ReadAll(res.Body)
	var actual = &CreatedResponse{}
	json.Unmarshal(resBodyByte, &actual)

	expect := new(CreatedResponse)
	expect.Id = uuid.New()

	if !cmp.Equal(actual, expect) {
		t.Errorf("expected %s, actual %s", expect, actual)		
	}
}