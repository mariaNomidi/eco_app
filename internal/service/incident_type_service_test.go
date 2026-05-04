package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/mariaNomidi/eco-app/internal/domain"
)

type mockIncidentTypeRepo struct {
	data []domain.IncidentType
	err  error
}

func (f *mockIncidentTypeRepo) GetAll(ctx context.Context) ([]domain.IncidentType, error) {
	return f.data, f.err
}

func TestGetAllIncidentTypes_Success(t *testing.T) {
	repo := &mockIncidentTypeRepo{
		data: []domain.IncidentType{
			{ID: 1, Name: "wildfire"},
			{ID: 2, Name: "injured_animal"},
		},
	}

	service := NewIncidentTypeService(repo)

	result, err := service.GetAll(context.Background())

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}
}

func TestGetAllIncidentTypes_Error(t *testing.T) {
	repo := &mockIncidentTypeRepo{
		err: fmt.Errorf("db error"),
	}

	service := NewIncidentTypeService(repo)

	_, err := service.GetAll(context.Background())

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
