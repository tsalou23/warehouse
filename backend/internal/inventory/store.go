package inventory

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

func (s *Store) GetAll(ctx context.Context) ([]Inventory, error) {
	rows, err := s.db.Query(ctx, "SELECT id,location_id,variant_id,quantity,updated_at FROM inventory")
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	var inventory []Inventory
	for rows.Next() {
		var w Inventory
		rows.Scan(&w.Id, &w.LocationId, &w.VariantId, &w.Quantity, &w.UpdatedAt)
		inventory = append(inventory, w)

	}
	return inventory, nil

}
