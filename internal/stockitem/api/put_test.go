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
	bforeName := uuid.NewString()
	afterName := uuid.NewString()

	// When
	postReqBody := &oapicodegen.PostStockItemJSONBody{
		Name: bforeName,
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
	putReqBody := &oapicodegen.PutStockItemJSONBody{
		Name: afterName,
	}
	putReqBodyJson, _ := json.Marshal(putReqBody)
	putReq, newReqErr := http.NewRequest(
		http.MethodPut,
		"http://localhost:3000/stock/items/"+postResBody.Id.String(),
		bytes.NewBuffer(putReqBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	putReq.Header.Set("Content-Type", "application/json")
	putRes, reqErr := client.Do(putReq)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer putRes.Body.Close()

	putResBodyByte, _ := io.ReadAll(putRes.Body)
	putResBody := &oapicodegen.Created{}
	json.Unmarshal(putResBodyByte, &putResBody)

	// Then
	if putRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, putRes.StatusCode)
	}
}