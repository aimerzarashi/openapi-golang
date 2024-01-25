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

func TestPutSuccess(t *testing.T) {

	// Setup
	client := http.Client{}
	stockItemName := uuid.NewString()
	updatedStockItemName := uuid.NewString()

	// Given
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

	if postResponse.StatusCode != http.StatusCreated {
		t.Fatal(reqErr)
	}

	if postResponseBody.Id == uuid.Nil {
		t.Fatal(reqErr)
	}

	// When
	putRequestBody := &oapicodegen.PutStockItemJSONBody{
		Name: updatedStockItemName,
	}
	putRequestBodyJson, _ := json.Marshal(putRequestBody)
	putRequest, newReqErr := http.NewRequest(
		http.MethodPut,
		"http://localhost:3000/stock/items/"+postResponseBody.Id.String(),
		bytes.NewBuffer(putRequestBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	putRequest.Header.Set("Content-Type", "application/json")
	putResponse, reqErr := client.Do(putRequest)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer putResponse.Body.Close()
	putResponseBodyByte, _ := io.ReadAll(putResponse.Body)
	putResponseBody := &oapicodegen.Created{}
	json.Unmarshal(putResponseBodyByte, &putResponseBody)

	// Then
	if putResponse.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, putResponse.StatusCode)
	}
}