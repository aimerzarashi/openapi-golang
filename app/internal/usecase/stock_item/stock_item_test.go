package stock_item

import (
	"bytes"
	"testing"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"

	"openapi/internal/presentation/stock_item_api"
)

// http://localhost:3000/ にGETでアクセスし、戻り値を検証する
func TestPostStockItem(t *testing.T) {
	request := stock_item_api.PostStockItemJSONBody{
		Name: uuid.NewString(),
	}
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
	var actual = &stock_item_api.CreatedResponse{}
	json.Unmarshal(resBodyByte, &actual)

	if actual.Id == uuid.Nil {
		t.Errorf("want not nil, got nil")
	}
}