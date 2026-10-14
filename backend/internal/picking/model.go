package picking

import "time"

type PickingAttempt struct {
	Id                int       `json:"id"`
	ItemId            int       `json:"item_id"`
	LocationId        int       `json:"location_id"`
	RequestedQuantity int       `json:"requested_quantity"`
	PickedQuantity    int       `json:"picked_quantity"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}