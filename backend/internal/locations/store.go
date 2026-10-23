package locations

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

func (s *Store) GetLocationsByWarehouseId(ctx context.Context, warehouseId int) ([]Location, error) {
	rows, err := s.db.Query(ctx, "SELECT id,warehouse_id,code,type,sort_order,created_at FROM locations WHERE warehouse_id =$1 ORDER BY sort_order", warehouseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var locations []Location
	for rows.Next() {
		var l Location
		rows.Scan(&l.Id, &l.WarehouseId, &l.Code, &l.Type, &l.SortOrder, &l.CreatedAt)
		locations = append(locations, l)

	}
	return locations, nil
}
