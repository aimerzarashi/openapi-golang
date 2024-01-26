package usecase

import (
	"testing"

	"github.com/google/uuid"
)


func TestCreateSuccess(t *testing.T) {
	// Given
	reqDto := &CreateRequestDto{
		Name: uuid.NewString(),
	}

	// When	
	resDto, err := Create(reqDto)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDto.Id == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, resDto.Id)
	}

}