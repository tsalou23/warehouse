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

func (s *Store) GetAllProductVariants(ctx context.Context, productId int) ([]ProductVariant, error) {
	rows, err := s.db.Query(ctx, "SELECT id,product_id,size_id,sku,created_at FROM product_variants WHERE product_id=$1", productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var variants []ProductVariant
	for rows.Next() {
		var w ProductVariant
		rows.Scan(&w.Id, &w.ProductId, &w.SizeId, &w.Sku, &w.CreatedAt)
		variants = append(variants, w)
	}
	return variants, nil
}

func (s *Store) GetVariantBySKU(ctx context.Context, sku string) (*ProductVariant, error) {
	var v ProductVariant
	err := s.db.QueryRow(ctx, "SELECT id,product_id,size_id,sku,created_at FROM product_variants WHERE sku=$1", sku).
		Scan(&v.Id, &v.ProductId, &v.SizeId, &v.Sku, &v.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
