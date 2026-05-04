package service

import (
	"context"

	"github.com/mariaNomidi/eco-app/internal/domain"
	"github.com/mariaNomidi/eco-app/internal/repository"
)

type IncidentTypeService struct {
	repo repository.IncidentTypeRepoInt
}

func NewIncidentTypeService(repo repository.IncidentTypeRepoInt) *IncidentTypeService {
	return &IncidentTypeService{repo: repo}
}

func (s *IncidentTypeService) GetAll(ctx context.Context) ([]domain.IncidentType, error) {
	return s.repo.GetAll(ctx)
}
