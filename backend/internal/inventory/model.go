package inventory

import "time"

type Inventory struct {
	Id         int       `json:"id"`
	LocationId int       `json:"location_id"`
	SizeId     int       `json:"size_id"`
	Quantity   int       `json:"quantity"`
	UpdatedAt  time.Time `json:"updated_at"`
}
