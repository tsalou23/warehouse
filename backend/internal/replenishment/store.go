package replenishment

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{db: db}
}

func (s *Store) GetAll(ctx context.Context) ([]ReplenishmentRequest, error) {
	rows, err := s.db.Query(ctx, "SELECT id,status,created_at,updated_at FROM replenishment_requests ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var replenishmentRequest []ReplenishmentRequest
	for rows.Next() {
		var v ReplenishmentRequest
		rows.Scan(&v.Id, &v.Status, &v.CreatedAt, &v.UpdatedAt)
		replenishmentRequest = append(replenishmentRequest, v)

	}
	return replenishmentRequest, nil

}

func (s *Store) Create(ctx context.Context, items []ReplenishmentItem) (*ReplenishmentRequest, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	var req ReplenishmentRequest
	err = tx.QueryRow(ctx, "INSERT INTO replenishment_requests(status,created_at,updated_at) VALUES($1,NOW(),NOW()) RETURNING id,status,created_at,updated_at", "created").Scan(&req.Id, &req.Status, &req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		_, err = tx.Exec(ctx, "INSERT INTO replenishment_items(request_id,variant_id,required_quantity,total_picked_quantity,created_at,updated_at) VALUES ($1,$2,$3,0,NOW(),NOW())", req.Id, item.VariantId, item.RequiredQuantity)
		if err != nil {
			return nil, err
		}
	}
	tx.Commit(ctx)
	return &req, nil
}
func (s *Store) AddItems(ctx context.Context, requestId int, items []ReplenishmentItem) error {
	for _, item := range items {
		_, err := s.db.Exec(ctx, "INSERT INTO replenishment_items(request_id,variant_id,required_quantity,total_picked_quantity,created_at,updated_at) VALUES ($1,$2,$3,0,NOW(),NOW())", requestId, item.VariantId, item.RequiredQuantity)
		if err != nil {
			return err
		}
	}
	return nil
}
