package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	store := warehouse.NewStore(conn)
	handler := warehouse.NewHandler(store)
	r.Get("/warehouses", handler.GetAll)

	http.ListenAndServe(":8080", r)

}
