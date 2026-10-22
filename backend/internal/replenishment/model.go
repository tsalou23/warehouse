package replenishment

import "time"

type ReplenishmentRequest struct {
	Id        int       `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReplenishmentItem struct {
	Id               int       `json:"id"`
	RequestId        int       `json:"request_id"`
	VariantId            int       `json:"variant_id"`
	RequiredQuantity int       `json:"required_quantity"`
	TotalPickedQuantity int       `json:"total_picked_quantity"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
