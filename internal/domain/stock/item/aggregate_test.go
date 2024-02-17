package item

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewAggregate(t *testing.T) {
	t.Parallel()

	// Setup
	id := uuid.New()

	type args struct {
		id   Id
		name Name
	}
	tests := []struct {
		name string
		args args
		want *Aggregate
	}{
		{
			name: "success",
			args: args{
				id: Id{
					value: id,
				},
				name: Name{
					value: "test",
				},
			},
			want: &Aggregate{
				Id: Id{
					value: id,
				},
				Name: Name{
					value: "test",
				},
				deleted: false,
			},
		},
	}
	// Run
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewAggregate(tt.args.id, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestoreAggregate(t *testing.T) {
	t.Parallel()

	id := uuid.New()

	type args struct {
		id      Id
		name    Name
		deleted bool
	}
	tests := []struct {
		name string
		args args
		want *Aggregate
	}{
		{
			name: "success",
			args: args{
				id: Id{
					value: id,
				},
				name: Name{
					value: "test",
				},
				deleted: false,
			},
			want: &Aggregate{
				Id: Id{
					value: id,
				},
				Name: Name{
					value: "test",
				},
				deleted: false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := RestoreAggregate(tt.args.id, tt.args.name, tt.args.deleted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregate_IsDeleted(t *testing.T) {
	t.Parallel()

	type args struct {
		Id      Id
		Name    Name
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
				deleted: false,
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				deleted: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := Aggregate{
				Id:      tt.args.Id,
				Name:    tt.args.Name,
				deleted: tt.args.deleted,
			}
			if got := a.IsDeleted(); got != tt.want {
				t.Errorf("Aggregate.IsDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregate_Delete(t *testing.T) {
	t.Parallel()

	// Given
	type args struct {
		Id      Id
		Name    Name
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
				deleted: false,
			},
			want: true,
		},
		{
			name: "no change",
			args: args{
				deleted: true,
			},
			want: true,
		},
	}

	// When & Then
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := &Aggregate{
				Id:      tt.args.Id,
				Name:    tt.args.Name,
				deleted: tt.args.deleted,
			}
			a.Delete()
		})
	}
}
