package picking

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

func (s *Store) Create(ctx context.Context, attempt PickingAttempt) (*PickingAttempt, error) {
	var a PickingAttempt
	err := s.db.QueryRow(ctx, `INSERT INTO picking_attempts
	(item_id,location_id,requested_quantity,picked_quantity,status,created_at,updated_at)
	VALUES($1,$2,$3,$4,$5,NOW(),NOW())
	RETURNING id,item_id,location_id,requested_quantity,picked_quantity,status,created_at,updated_at`,
		attempt.ItemId, attempt.LocationId, attempt.RequestedQuantity, attempt.PickedQuantity, attempt.Status,
	).Scan(&a.Id, &a.ItemId, &a.LocationId, &a.RequestedQuantity, &a.PickedQuantity, &a.Status, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
func (s *Store) UpdateStatus(ctx context.Context, id, pickedQuantity int, status string) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var itemId int
	var locationId int
	err = tx.QueryRow(ctx, "SELECT item_id,location_id FROM picking_attempts WHERE id=$1", id).Scan(&itemId, &locationId)
	if err != nil {
		return err
	}
	var variantId int
	err = tx.QueryRow(ctx, "SELECT variant_id FROM replenishment_items WHERE id=$1", itemId).Scan(&variantId)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "UPDATE picking_attempts SET status=$1,picked_quantity=$2,updated_at=NOW() WHERE id=$3", status, pickedQuantity, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE replenishment_items SET total_picked_quantity=total_picked_quantity+$1 ,updated_at=NOW() WHERE id =$2", pickedQuantity, itemId)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "UPDATE inventory SET quantity=quantity-$1,updated_at=NOW() WHERE location_id=$2 AND variant_id=$3", pickedQuantity, locationId, variantId)
	if err != nil {
		return err
	}
	tx.Commit(ctx)
	return nil

}
