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

func TestPostCreated(t *testing.T) {
	t.Parallel()

	// When
	reqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: "test",
	}

	r := NewRequest("/stock/locations", reqBody)
	
	err := locations.Api.PostStockLocation(locations.Api{}, r.context)
	if err != nil {
		t.Fatal(err)
	}
	defer r.recorder.Result().Body.Close()

	// Then
	if r.recorder.Code != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, r.recorder.Code)
	}

	postResBody, err := Response[oapicodegen.Created](r.recorder.Result())
	if err != nil {
		t.Fatal(err)
	}

	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}
}

func TestPostBadRequest(t *testing.T) {
	// Setup
	rh := RequestHelper{
		client: &http.Client{},
	}
	rch := ResponseConvertHelper{}
	
	zeroLenName := ""
	overLenName := strings.Repeat("a", 101)

	// When
	postResZeroLen, err := rh.Post(
		&oapicodegen.PostStockLocationJSONRequestBody{
			Name: zeroLenName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postResZeroLen.Body.Close()

	postResOverLen, err := rh.Post(
		&oapicodegen.PostStockLocationJSONRequestBody{
			Name: overLenName,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postResOverLen.Body.Close()

	// Then
	if postResZeroLen.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postResZeroLen.StatusCode)
	}	

	postResZeroLenBody, err := rch.AsBadRequest(postResZeroLen)
	if err != nil {
		t.Fatal(err)
	}
	if postResZeroLenBody.Message == "" {
		t.Errorf("expected empty, actual %s", postResZeroLenBody.Message)
	}
	
	if postResOverLen.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postResOverLen.StatusCode)
	}	

	postResOverLenBody, err := rch.AsBadRequest(postResOverLen)
	if err != nil {
		t.Fatal(err)
	}
	if postResOverLenBody.Message== "" {
		t.Errorf("expected empty, actual %s", postResOverLenBody.Message)
	}
}