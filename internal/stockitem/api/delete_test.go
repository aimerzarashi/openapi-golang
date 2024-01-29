package api

import (
	"openapi/internal/infra/env"
	oapicodegen "openapi/internal/infra/oapicodegen/stockitem"
	"testing"

	"bytes"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

func TestDeleteOk(t *testing.T) {

	// Setup
	client := http.Client{}
	name := uuid.NewString()

	// Given
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

	if postRes.StatusCode != http.StatusCreated {
		t.Fatal(err)
	}

	postResBodyByte, _ := io.ReadAll(postRes.Body)
	postResBody := &oapicodegen.Created{}
	json.Unmarshal(postResBodyByte, &postResBody)
	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

	// When
	deleteReq, err := http.NewRequest(
		http.MethodDelete,
		env.GetServiceUrl()+"/stock/items/"+postResBody.Id.String(),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteRes, err := client.Do(deleteReq)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	deleteResBodyByte, _ := io.ReadAll(deleteRes.Body)
	deleteResBody := &oapicodegen.Created{}
	json.Unmarshal(deleteResBodyByte, &deleteResBody)

	// Then
	if deleteRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, deleteRes.StatusCode)
	}
}

func TestDeleteNotFound(t *testing.T) {
	// Setup
	client := http.Client{}

	deleteReq, err := http.NewRequest(
		http.MethodDelete,
		env.GetServiceUrl()+"/stock/items/"+uuid.NewString(),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteRes, err := client.Do(deleteReq)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	deleteResBodyByte, _ := io.ReadAll(deleteRes.Body)
	deleteResBody := &oapicodegen.Created{}
	json.Unmarshal(deleteResBodyByte, &deleteResBody)

	// Then
	if deleteRes.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, deleteRes.StatusCode)
	}
}