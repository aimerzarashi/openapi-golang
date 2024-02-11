package locations_test

import (
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
	"openapi/internal/ui/stock/locations"
	"strings"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestPutOk2(t *testing.T) {
	t.Parallel()

	// Given
	beforeReqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: "test",
	}

	b := NewRequest("/stock/locations", beforeReqBody)
	
	if err := locations.Api.PostStockLocation(locations.Api{}, b.context); err != nil {
		t.Fatal(err)
	}
	defer b.recorder.Result().Body.Close()

	if b.recorder.Code != http.StatusCreated {
		t.Fatal(b.recorder.Code)
	}

	postReqBody, err := Response[oapicodegen.Created](b.recorder.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	afterReqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: "test",
	}

	a := NewRequest("/stock/locations/" + postReqBody.Id.String(), afterReqBody)
	
	if err := locations.Api.PutStockLocation(locations.Api{}, a.context, postReqBody.Id); err != nil {
		t.Fatal(err)
	}
	defer a.recorder.Result().Body.Close()

	// Then
	if a.recorder.Code != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, a.recorder.Code)
	}
}

func TestPutOk(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}
	
	bforeName := "test"
	afterName := "newTest"

	// Given
	postRes, err := rh.Post(
		&oapicodegen.PostStockLocationJSONRequestBody{
			Name: bforeName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	if postRes.StatusCode != http.StatusCreated {
		t.Fatalf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBody, err := rch.AsCreated(postRes)
	if err != nil {
		t.Fatal(err)
	}

	// When
	putRes, err := rh.Put(
		postResBody.Id,
		&oapicodegen.PutStockLocationJSONRequestBody{
			Name: afterName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()
	
	// Then
	if putRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, putRes.StatusCode)
	}
}

func TestPutNotFound(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}

	name := uuid.NewString()

	putRes, err := rh.Put(
		uuid.New(),
		&oapicodegen.PutStockLocationJSONRequestBody{
			Name: name,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()

	// Then
	if putRes.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, putRes.StatusCode)
	}
}

func TestPutBadRequest(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}
	
	zeroLenName := ""
	overLenName := strings.Repeat("a", 101)

	// Given
	postRes, err := rh.Post(
		&oapicodegen.PostStockLocationJSONRequestBody{
			Name: uuid.NewString(),
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	if postRes.StatusCode != http.StatusCreated {
		t.Fatalf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBody, err := rch.AsCreated(postRes)
	if err != nil {
		t.Fatal(err)
	}

	// When
	putResZeroLen, err := rh.Put(
		postResBody.Id,
		&oapicodegen.PutStockLocationJSONRequestBody{
			Name: zeroLenName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer putResZeroLen.Body.Close()

	putResOverLen, err := rh.Put(
		postResBody.Id,
		&oapicodegen.PutStockLocationJSONRequestBody{
			Name: overLenName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer putResOverLen.Body.Close()

	// Then
	if putResZeroLen.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, putResZeroLen.StatusCode)
	}

	putResBodyZeroLen, err := rch.AsBadRequest(putResZeroLen)
	if err != nil {
		t.Fatal(err)
	}
	if putResBodyZeroLen.Message == "" {
		t.Errorf("expected not empty, actual empty")
	}
	
	if putResOverLen.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, putResOverLen.StatusCode)
	}

	putResBodyOverLen, err := rch.AsBadRequest(putResOverLen)
	if err != nil {
		t.Fatal(err)
	}
	if putResBodyOverLen.Message == "" {
		t.Errorf("expected not empty, actual empty")
	}
}