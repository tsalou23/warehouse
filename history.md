# Project History — Warehouse Operations Platform

## Project Structure
```
inditex/
  backend/
    cmd/api/main.go
    internal/
      warehouse/
        model.go
        store.go
        handler.go
      products/
        model.go
        store.go
        handler.go
      replenishment/   ← not started yet
  database/
    schema.sql
  frontend/            ← not started yet
  .env
  .gitignore
  docker-compose.yml
```

---

## Phase 1 — Domain Design

### Domain Entities agreed
- Product (SKU, name)
- ProductVariant (product + size)
- Warehouse (name, type)
- Location (warehouse, code, sort_order, type)
- Inventory (location + variant + quantity)
- ReplenishmentRequest (status)
- ReplenishmentItem (request + variant + required/picked quantity)
- WarehouseAllocation (item + warehouse + allocated quantity)
- PickingAttempt (item + location + requested/picked quantity + status)

### Business Rules agreed
- picked_quantity can never exceed required_quantity
- remaining_quantity = required - picked (NOT stored, always computed)
- Inventory never goes below 0
- ReplenishmentRequest auto-completes when all items have remaining = 0
- One replenishment = one worker (V1)
- Inventory decremented only on confirmed pick (no reservations in V1)
- AllocationStrategy is pluggable (interface) — V1 uses simple placeholder

### State Machines
ReplenishmentRequest: created → in_progress → completed / cancelled
PickingAttempt: pending → found / not_found / partially_found

---

## Phase 2 — Database Schema

### Tables created (database/schema.sql)
- products
- warehouses
- locations
- product_variants
- inventory
- replenishment_requests
- replenishment_items
- warehouse_allocations
- picking_attempts

### Key decisions
- remaining_quantity is computed (required - picked), NOT stored
- status fields use VARCHAR + CHECK constraints (not ENUM) for extensibility
- All tables have created_at, mutable tables also have updated_at
- Inventory has CHECK quantity >= 0
- Composite unique constraints where needed (e.g. location_id + variant_id)

---

## Phase 3 — Development Environment

### Setup completed
- Git repo initialized
- GitHub repo: https://github.com/tsalou23/warehouse.git
- Docker Desktop installed
- docker-compose.yml created (PostgreSQL 16)
- .env file created (DATABASE_URL)
- .gitignore created
- Go module initialized: github.com/tsalou23/warehouse

### Dependencies installed
- github.com/go-chi/chi/v5 (routing)
- github.com/jackc/pgx/v5 (PostgreSQL driver)
- github.com/joho/godotenv (env vars)

### Docker commands
```bash
docker compose up -d        # start PostgreSQL
docker compose down         # stop PostgreSQL
docker exec -i inditex-db-1 psql -U warehouse -d warehouse < database/schema.sql  # run schema
```

---

## Phase 4 — Backend

### Architecture: Repository Pattern + Layered Architecture
Each domain module has:
- model.go  → Go structs (translation of SQL table)
- store.go  → SQL queries (only place that talks to DB)
- handler.go → HTTP handlers (only place that handles HTTP)
- main.go connects everything

### Pattern used everywhere
```go
// 1. Define the struct
type Store struct { db *pgx.Conn }

// 2. Constructor
func NewStore(db *pgx.Conn) *Store { return &Store{db: db} }

// 3. Methods
func (s *Store) GetAll(ctx context.Context) ([]Warehouse, error) { ... }
```

### Completed endpoints
- GET /health → returns "OK"
- GET /warehouses → returns all warehouses from DB
- GET /products → returns all products from DB (needs wiring in main.go)
- GET /products/{id}/variants → returns variants of a product (needs wiring in main.go)

---

## Current State of main.go

```go
// warehouse wired correctly
store := warehouse.NewStore(conn)      // ← needs rename to warehouseStore
handler := warehouse.NewHandler(store) // ← needs rename to warehouseHandler

// products added but has naming inconsistency
productsStore := products.NewStore(conn)
productsHandler := products.NewHandler(productStore) // ← bug: productStore vs productsStore
r.Get("/products", productHandler.GetAll)
r.Get("/products/{id}/variants", productHandler.GetProductVariants)
```

### Issues to fix in main.go
1. Rename warehouse store/handler to warehouseStore/warehouseHandler for consistency
2. Fix variable name bug: productsStore vs productStore
3. Add products import

---

## Next Steps
1. Fix main.go naming issues
2. Add products import to main.go
3. Test GET /products and GET /products/{id}/variants
4. Add mock data for products and product_variants
5. Start replenishment module
6. Start frontend

---

## Concepts learned so far
- SQL schema design (tables, constraints, foreign keys, indexes)
- Docker and Docker Compose
- Go modules (go.mod, go.sum)
- HTTP basics (methods, routes, handlers)
- REST API design
- Repository Pattern
- Go structs, methods, pointers (&, *)
- JSON tags
- pgx (PostgreSQL driver)
- context.Background() vs r.Context()
- SQL injection prevention ($1 parameters)
- camelCase vs PascalCase vs snake_case conventions
- Separation of concerns (model/store/handler)
