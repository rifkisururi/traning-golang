-- Migration: Create kategori table
-- Created: 2026-01-31

CREATE TABLE IF NOT EXISTS kategori (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    deskripsi TEXT
);

-- Insert initial data
INSERT INTO kategori (id, nama, deskripsi) VALUES
(1, 'Makanan', 'Kategori untuk produk makanan'),
(2, 'Minuman', 'Kategori untuk produk minuman'),
(3, 'Bumbu Dapur', 'Kategori untuk bumbu dan rempah');
