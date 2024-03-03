package item_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/aimerzarashi/timeslice"
	"github.com/google/uuid"

	"openapi/internal/domain/sell/item"
	"openapi/internal/domain/sell/item/value"
)

func TestNewAggregate(t *testing.T) {
	t.Parallel()

	// Given
	id, err := item.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := value.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		id   item.Id
		name value.Name
	}

	type want struct {
		id      item.Id
		name    value.Name
		deleted bool
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				id:   id,
				name: name,
			},
			want: want{
				id:      id,
				name:    name,
				deleted: false,
			},
		},
	}

	// Run
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got, err := item.NewAggregate(tt.args.id, tt.args.name)

			// Then
			if err != nil {
				t.Errorf("NewAggregate() error = %v", err)
				return
			}

			if got.Id 				 != tt.want.id || 
				 got.Name 			!= tt.want.name || 
				 got.IsDeleted() != tt.want.deleted {
				t.Errorf("NewAggregate() = %v, want %v", got, tt.want)
			}

			price, err := value.NewPrice(100, "JPY")
			if err != nil {
				t.Fatal(err)
			}
			item, err := timeslice.NewItem(price, time.Now(), time.Now().Add(time.Minute*10))
			if err != nil {
				t.Fatal(err)
			}
			prices, err := got.Prices.Add(item)
			if err != nil {
				t.Fatal(err)
			}
			for _, p := range prices.Items() {
				fmt.Printf("%+v, %+v, %+v\n", p.Value(), p.StartAt(), p.EndAt())				
			}

		})
	}
}

func TestRestoreAggregate(t *testing.T) {
	t.Parallel()

	// Given
	id, err := item.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := value.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		id      item.Id
		name    value.Name
		deleted bool
	}
	type want struct {
		Id      item.Id
		Name    value.Name
		deleted bool
	}
	tests := []struct {
		name string
		args args
		want *want
	}{
		{
			name: "active",
			args: args{
				id:      id,
				name:    name,
				deleted: false,
			},
			want: &want{
				Id:      id,
				Name:    name,
				deleted: false,
			},
		},
		{
			name: "deleted",
			args: args{
				id:      id,
				name:    name,
				deleted: true,
			},
			want: &want{
				Id:      id,
				Name:    name,
				deleted: true,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			got := item.RestoreAggregate(tt.args.id, tt.args.name, tt.args.deleted)

			// Then
			if !reflect.DeepEqual(got.Id, tt.want.Id) {
				t.Errorf("RestoreAggregate() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("RestoreAggregate() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.IsDeleted(), tt.want.deleted) {
				t.Errorf("RestoreAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregate_Delete(t *testing.T) {
	t.Parallel()

	// Setup
	id, err := item.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := value.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		Id      item.Id
		Name    value.Name
		deleted bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				Id:      id,
				Name:    name,
				deleted: false,
			},
			want: true,
		},
		{
			name: "no change",
			args: args{
				Id:      id,
				Name:    name,
				deleted: true,
			},
			want: true,
		},
	}

	// Run
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			a := item.RestoreAggregate(tt.args.Id, tt.args.Name, tt.args.deleted)

			// When
			a.Delete()

			// Then
			if got := a.IsDeleted(); got != tt.want {
				t.Errorf("Aggregate.IsDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
