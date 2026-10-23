package locations

import "time"

type Location struct {
	Id          int       `json:"id"`
	WarehouseId int       `json:"warehouse_id"`
	Code        int       `json:"code"`
	Type        string    `json:"type"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
}
