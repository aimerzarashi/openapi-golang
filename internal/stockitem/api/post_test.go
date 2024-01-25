package api

import (
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
	"testing"

	"bytes"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

func TestPostSuccess(t *testing.T) {

	// Setup
	client := http.Client{}
	stockItemName := uuid.NewString()

	// When
	postRequestBody := &oapicodegen.PostStockItemJSONBody{
		Name: stockItemName,
	}
	postRequestBodyJson, _ := json.Marshal(postRequestBody)
	postRequest, newReqErr := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(postRequestBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	postRequest.Header.Set("Content-Type", "application/json")
	postResponse, reqErr := client.Do(postRequest)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer postResponse.Body.Close()
	postResponseBodyByte, _ := io.ReadAll(postResponse.Body)
	postResponseBody := &oapicodegen.Created{}
	json.Unmarshal(postResponseBodyByte, &postResponseBody)

	// Then
	if postResponse.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, postResponse.StatusCode)
	}

	if postResponseBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

}
