CREATE TABLE products(
    id  SERIAL PRIMARY KEY,
    sku TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL

);
CREATE TABLE warehouses(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    type VARCHAR NOT NULL CHECK(type IN('hanging','folded')),
    created_at TIMESTAMP NOT NULL

);
CREATE TABLE locations(
    id SERIAL PRIMARY KEY,
    warehouse_id INTEGER NOT NULL REFERENCES warehouses(id),
    code INT NOT NULL UNIQUE,
    type VARCHAR NOT NULL CHECK(type IN('hanging','folded')),
    sort_order INT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
CREATE TABLE product_variants(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id),
    size TEXT NOT NULL CHECK (size IN('xs','s','m','l','xl')),
    created_at TIMESTAMP NOT NULL,
    UNIQUE(product_id,size)
);
CREATE TABLE inventory(
    id SERIAL PRIMARY KEY,
    location_id INTEGER NOT NULL REFERENCES locations(id),
    variant_id INTEGER NOT NULL REFERENCES product_variants(id),
    quantity INTEGER NOT NULL CHECK (quantity>=0),
    UNIQUE(location_id,variant_id),
    updated_at TIMESTAMP NOT NULL

);
CREATE TABLE replenishment_requests(
    id SERIAL PRIMARY KEY,
    status VARCHAR NOT NULL CHECK (status IN('created','in_progress','completed','cancelled')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE TABLE replenishment_items(
    id SERIAL PRIMARY KEY,
    request_id INTEGER NOT NULL REFERENCES replenishment_requests(id),
    variant_id INTEGER NOT NULL REFERENCES product_variants(id),
    required_quantity INTEGER NOT NULL CHECK( required_quantity>0),
    picked_quantity INTEGER NOT NULL CHECK(picked_quantity>=0),
    UNIQUE (request_id,variant_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE TABLE warehouse_allocations(
    id SERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL REFERENCES replenishment_items(id),
    warehouse_id INTEGER NOT NULL REFERENCES warehouses(id),
    allocated_quantity INTEGER NOT NULL CHECK (allocated_quantity>0),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (item_id,warehouse_id)
);

CREATE TABLE picking_attempts(
    id SERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL REFERENCES replenishment_items(id),
    location_id INTEGER NOT NULL REFERENCES locations(id),
    requested_quantity INTEGER NOT NULL CHECK (requested_quantity > 0),
    picked_quantity INTEGER NOT NULL CHECK (picked_quantity >= 0),
    status VARCHAR NOT NULL CHECK (status IN('pending','found','not_found','partially_found')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
