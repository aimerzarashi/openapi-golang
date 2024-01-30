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

func TestPutOk(t *testing.T) {

	// Setup
	client := http.Client{}
	bforeName := uuid.NewString()
	afterName := uuid.NewString()

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: bforeName,
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
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: afterName,
	}
	putReqBodyJson, _ := json.Marshal(putReqBody)
	putReq, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+postResBody.Id.String(),
		bytes.NewBuffer(putReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}
	putReq.Header.Set("Content-Type", "application/json")
	putRes, err := client.Do(putReq)
	if err != nil {
		t.Fatal(err)
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


func TestPutNotFound(t *testing.T) {
	// Setup
	client := http.Client{}
	name := uuid.NewString()

	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: name,
	}
	putReqBodyJson, _ := json.Marshal(putReqBody)
	putReq, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+uuid.NewString(),
		bytes.NewBuffer(putReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}
	putReq.Header.Set("Content-Type", "application/json")
	putRes, err := client.Do(putReq)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()

	putResBodyByte, _ := io.ReadAll(putRes.Body)
	putResBody := &oapicodegen.Created{}
	json.Unmarshal(putResBodyByte, &putResBody)

	// Then
	if putRes.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, putRes.StatusCode)
	}
}

func TestPutBadRequest1(t *testing.T) {
	// Setup
	client := http.Client{}
	bforeName := uuid.NewString()
	afterName := ""

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: bforeName,
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
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: afterName,
	}
	putReqBodyJson, _ := json.Marshal(putReqBody)
	putReq, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+postResBody.Id.String(),
		bytes.NewBuffer(putReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}
	putReq.Header.Set("Content-Type", "application/json")
	putRes, err := client.Do(putReq)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()

	putResBodyByte, _ := io.ReadAll(putRes.Body)
	putResBody := &oapicodegen.Created{}
	json.Unmarshal(putResBodyByte, &putResBody)

	// Then
	if putRes.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, putRes.StatusCode)
	}
}

func TestPutBadRequest2(t *testing.T) {
	// Setup
	client := http.Client{}
	bforeName := uuid.NewString()
	afterName := strings.Repeat("a", 101)

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: bforeName,
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
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: afterName,
	}
	putReqBodyJson, _ := json.Marshal(putReqBody)
	putReq, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+postResBody.Id.String(),
		bytes.NewBuffer(putReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}
	putReq.Header.Set("Content-Type", "application/json")
	putRes, err := client.Do(putReq)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()

	putResBodyByte, _ := io.ReadAll(putRes.Body)
	putResBody := &oapicodegen.Created{}
	json.Unmarshal(putResBodyByte, &putResBody)

	// Then
	if putRes.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, putRes.StatusCode)
	}
}