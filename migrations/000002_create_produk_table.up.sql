-- Migration: Create produk table
-- Created: 2026-01-31

CREATE TABLE IF NOT EXISTS produk (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    harga INTEGER NOT NULL,
    stok INTEGER NOT NULL DEFAULT 0
);

-- Insert initial data
INSERT INTO produk (id, nama, harga, stok) VALUES
(1, 'Indomie Godog', 3500, 10),
(2, 'Vit 1000ml', 3000, 40),
(3, 'kecap', 12000, 20);
