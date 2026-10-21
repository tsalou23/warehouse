-- Products
INSERT INTO products (sku, name, created_at) VALUES
('1234/567/890', 'Classic White Shirt', NOW()),
('2345/678/901', 'Slim Fit Jeans', NOW()),
('3456/789/012', 'Wool Blazer', NOW());

-- Product Sizes (sku = product_sku/01 for xs, /02 for s, /03 for m, /04 for l, /05 for xl)
INSERT INTO product_sizes (product_id, sku, size, created_at) VALUES
(1, '1234/567/890/01', 'xs', NOW()),
(1, '1234/567/890/02', 's',  NOW()),
(1, '1234/567/890/03', 'm',  NOW()),
(1, '1234/567/890/04', 'l',  NOW()),
(1, '1234/567/890/05', 'xl', NOW()),
(2, '2345/678/901/02', 's',  NOW()),
(2, '2345/678/901/03', 'm',  NOW()),
(2, '2345/678/901/04', 'l',  NOW()),
(2, '2345/678/901/05', 'xl', NOW()),
(3, '3456/789/012/02', 's',  NOW()),
(3, '3456/789/012/03', 'm',  NOW()),
(3, '3456/789/012/04', 'l',  NOW());

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
INSERT INTO inventory (location_id, size_id, quantity, updated_at) VALUES
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
INSERT INTO inventory (location_id, size_id, quantity, updated_at) VALUES
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
INSERT INTO replenishment_items (request_id, size_id, required_quantity, total_picked_quantity, created_at, updated_at) VALUES
(1, 3,  4, 0, NOW(), NOW()),
(1, 7,  3, 0, NOW(), NOW()),
(1, 11, 2, 0, NOW(), NOW());
