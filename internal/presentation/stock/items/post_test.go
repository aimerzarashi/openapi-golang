package items_test

import (
	"openapi/internal/infrastructure/env"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"strings"
	"testing"

	"bytes"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

func TestPostCreated(t *testing.T) {
	// Setup
	client := http.Client{}
	name := uuid.NewString()

	// When
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: name,
	}
	postReqBodyJson, _ := json.Marshal(postReqBody)
	postReq, err := http.NewRequest(
		http.MethodPost,
		env.GetServiceUrl()+"/stock/items",
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

func TestPostBadRequest1(t *testing.T) {
	// Setup
	client := http.Client{}
	name := strings.Repeat("a", 101)

	// When
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: name,
	}
	postReqBodyJson, _ := json.Marshal(postReqBody)
	postReq, err := http.NewRequest(
		http.MethodPost,
		env.GetServiceUrl()+"/stock/items",
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
	if postRes.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postRes.StatusCode)
	}	
}

func TestPostBadRequest2(t *testing.T) {
	// Setup
	client := http.Client{}
	name := ""

	// When
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: name,
	}
	postReqBodyJson, _ := json.Marshal(postReqBody)
	postReq, err := http.NewRequest(
		http.MethodPost,
		env.GetServiceUrl()+"/stock/items",
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
	if postRes.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postRes.StatusCode)
	}	
}