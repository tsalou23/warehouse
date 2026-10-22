-- Sizes lookup table
INSERT INTO sizes (label, sku_suffix, sort_order) VALUES
('xs', '01', 1),
('s',  '02', 2),
('m',  '03', 3),
('l',  '04', 4),
('xl', '05', 5);

-- Products
INSERT INTO products (sku, name, created_at) VALUES
('1234/567/890', 'Classic White Shirt', NOW()),
('2345/678/901', 'Slim Fit Jeans',      NOW()),
('3456/789/012', 'Wool Blazer',         NOW());

-- Product Variants
INSERT INTO product_variants (product_id, size_id, sku, created_at) VALUES
(1, 1, '1234/567/890/01', NOW()),
(1, 2, '1234/567/890/02', NOW()),
(1, 3, '1234/567/890/03', NOW()),
(1, 4, '1234/567/890/04', NOW()),
(1, 5, '1234/567/890/05', NOW()),
(2, 2, '2345/678/901/02', NOW()),
(2, 3, '2345/678/901/03', NOW()),
(2, 4, '2345/678/901/04', NOW()),
(2, 5, '2345/678/901/05', NOW()),
(3, 2, '3456/789/012/02', NOW()),
(3, 3, '3456/789/012/03', NOW()),
(3, 4, '3456/789/012/04', NOW());

-- Warehouses
INSERT INTO warehouses (name, type, created_at) VALUES
('Hanging Storage A', 'hanging', NOW()),
('Batch Storage B',   'folded',  NOW());

-- Locations (Hanging Storage A)
INSERT INTO locations (warehouse_id, code, type, sort_order, created_at) VALUES
(1, 1000, 'hanging', 1, NOW()),
(1, 1001, 'hanging', 2, NOW()),
(1, 1002, 'hanging', 3, NOW()),
(1, 1003, 'hanging', 4, NOW()),
(1, 1004, 'hanging', 5, NOW());

-- Locations (Batch Storage B)
INSERT INTO locations (warehouse_id, code, type, sort_order, created_at) VALUES
(2, 2000, 'folded', 1, NOW()),
(2, 2001, 'folded', 2, NOW()),
(2, 2002, 'folded', 3, NOW()),
(2, 2003, 'folded', 4, NOW()),
(2, 2004, 'folded', 5, NOW());

-- Inventory (Hanging Storage A)
INSERT INTO inventory (location_id, variant_id, quantity, updated_at) VALUES
(1, 3,  5, NOW()),
(1, 7,  3, NOW()),
(2, 4,  8, NOW()),
(2, 8,  6, NOW()),
(3, 5,  4, NOW()),
(3, 11, 2, NOW()),
(4, 3,  7, NOW()),
(4, 6,  5, NOW()),
(5, 4,  3, NOW()),
(5, 12, 4, NOW());

-- Inventory (Batch Storage B)
INSERT INTO inventory (location_id, variant_id, quantity, updated_at) VALUES
(6,  6,  10, NOW()),
(6,  10, 8,  NOW()),
(7,  7,  6,  NOW()),
(7,  11, 4,  NOW()),
(8,  8,  9,  NOW()),
(8,  12, 3,  NOW()),
(9,  3,  5,  NOW()),
(9,  6,  7,  NOW()),
(10, 4,  6,  NOW()),
(10, 7,  8,  NOW());

-- Replenishment Request
INSERT INTO replenishment_requests (status, created_at, updated_at) VALUES
('created', NOW(), NOW());

-- Replenishment Items
INSERT INTO replenishment_items (request_id, variant_id, required_quantity, total_picked_quantity, created_at, updated_at) VALUES
(1, 3,  4, 0, NOW(), NOW()),
(1, 7,  3, 0, NOW(), NOW()),
(1, 11, 2, 0, NOW(), NOW());
