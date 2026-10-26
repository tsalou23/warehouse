package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/tsalou23/warehouse/internal/inventory"
	"github.com/tsalou23/warehouse/internal/locations"
	"github.com/tsalou23/warehouse/internal/picking"
	"github.com/tsalou23/warehouse/internal/products"
	"github.com/tsalou23/warehouse/internal/replenishment"
	"github.com/tsalou23/warehouse/internal/warehouse"
)

func main() {
	fmt.Println("Warehouse API starting...")
	godotenv.Load("../.env")
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}
	fmt.Println("Database connected")
	defer conn.Close(context.Background())
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	warehouseStore := warehouse.NewStore(conn)
	warehouseHandler := warehouse.NewHandler(warehouseStore)
	r.Get("/warehouses", warehouseHandler.GetAll)

	productStore := products.NewStore(conn)
	productHandler := products.NewHandler(productStore)
	r.Get("/products", productHandler.GetAll)
	r.Get("/products/{id}/variants", productHandler.GetProductVariants)

	inventoryStore := inventory.NewStore(conn)
	inventoryHandler := inventory.NewHandler(inventoryStore)
	r.Get("/inventory", inventoryHandler.GetAll)
	replenishmentStore := replenishment.NewStore(conn)
	replenishmentHandler := replenishment.NewHandler(replenishmentStore)
	r.Get("/replenishment-requests", replenishmentHandler.GetAll)
	r.Post("/replenishment-requests", replenishmentHandler.Create)
	r.Post("/replenishment-requests/{id}/items", replenishmentHandler.AddItems)
	r.Get("/replenishment-requests/{id}/next-location", replenishmentHandler.NextLocation)

	pickingStore := picking.NewStore(conn)
	pickingHandler := picking.NewHandler(pickingStore)
	r.Post("/picking-attempts", pickingHandler.Create)
	r.Patch("/picking-attempts/{id}", pickingHandler.UpdateStatus)
	locationsStore := locations.NewStore(conn)
	locationsHandler := locations.NewHandler(locationsStore)
	r.Get("/warehouses/{id}/locations", locationsHandler.GetLocationsByWarehouseId)

	http.ListenAndServe(":8080", r)

}
