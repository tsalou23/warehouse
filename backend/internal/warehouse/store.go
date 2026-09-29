package warehouse

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

func (s *Store) GetAll(ctx context.Context) ([]Warehouse, error) {
	rows, err := s.db.Query(ctx, "SELECT id,name,type,created_at FROM warehouses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var warehouses []Warehouse
	for rows.Next() {
		var w Warehouse
		rows.Scan(&w.Id, &w.Name, &w.Type, &w.CreatedAt)
		warehouses = append(warehouses, w)
	}
	return warehouses, nil
}
