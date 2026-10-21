package products

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

func (s *Store) GetAll(ctx context.Context) ([]Product, error) {
	rows, err := s.db.Query(ctx, "SELECT id,sku,name,created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var w Product
		rows.Scan(&w.Id, &w.Sku, &w.Name, &w.CreatedAt)
		products = append(products, w)
	}
	return products, nil
}

func (s *Store) GetAllProductSizes(ctx context.Context, productId int) ([]ProductSize, error) {
	rows, err := s.db.Query(ctx, "SELECT id,product_id,sku,size,created_at FROM product_sizes WHERE product_id=$1", productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productSizes []ProductSize
	for rows.Next() {
		var w ProductSize
		rows.Scan(&w.Id, &w.ProductId, &w.Sku, &w.Size, &w.CreatedAt)
		productSizes = append(productSizes, w)
	}
	return productSizes, nil
}
