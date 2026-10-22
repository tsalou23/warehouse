package products

import "time"

type Product struct {
	Id        int       `json:"id"`
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductVariant struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	SizeId    int       `json:"size_id"`
	Sku       string    `json:"sku"`
	CreatedAt time.Time `json:"created_at"`
}
