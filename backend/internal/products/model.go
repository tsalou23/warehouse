package products

import "time"

type Product struct {
	Id        int       `json:"id"`
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductSize struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	Sku       string    `json:"sku"`
	Size      string    `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}
