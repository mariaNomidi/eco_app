package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mariaNomidi/eco-app/internal/domain"
)

type IncidentTypeRepo struct {
	db *pgxpool.Pool
}

type IncidentTypeRepoInt interface {
	GetAll(ctx context.Context) ([]domain.IncidentType, error)
}

func NewIncidentTypeRepo(db *pgxpool.Pool) *IncidentTypeRepo {
	return &IncidentTypeRepo{db: db}
}

func (r *IncidentTypeRepo) GetAll(ctx context.Context) ([]domain.IncidentType, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, name
        FROM incident_types
        ORDER BY id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.IncidentType

	for rows.Next() {
		var it domain.IncidentType
		if err := rows.Scan(&it.ID, &it.Name); err != nil {
			return nil, err
		}
		result = append(result, it)
	}

	return result, nil
}
