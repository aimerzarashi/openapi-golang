package usecase

import (
	"testing"

	"github.com/google/uuid"
)

func TestUpdateSuccess(t *testing.T) {
	// Given
	beforeName := uuid.NewString()
	afterName := uuid.NewString()

	reqCreateDto := &CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := Create(reqCreateDto)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqUpdateDto := &UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	resUpdateDto, err := Update(reqUpdateDto)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resUpdateDto == nil {
		t.Errorf("expected not empty, actual empty")
	}
}