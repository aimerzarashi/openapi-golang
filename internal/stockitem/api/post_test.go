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
	name := uuid.NewString()

	// When
	postReqBody := &oapicodegen.PostStockItemJSONBody{
		Name: name,
	}
	postReqBodyJson, _ := json.Marshal(postReqBody)
	postReq, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(postReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}
	postReq.Header.Set("Content-Type", "application/json")
	postRes, err := client.Do(postReq)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	// Then
	if postRes.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBodyByte, _ := io.ReadAll(postRes.Body)
	postResBody := &oapicodegen.Created{}
	json.Unmarshal(postResBodyByte, &postResBody)
	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

}
