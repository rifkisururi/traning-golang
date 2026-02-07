-- Migration: Add kategori_id column to produk table with foreign key
-- Created: 2026-01-31

-- Add kategori_id column to produk table
ALTER TABLE produk ADD COLUMN kategori_id INTEGER;

-- Add foreign key constraint
ALTER TABLE produk
ADD CONSTRAINT fk_produk_kategori
FOREIGN KEY (kategori_id)
REFERENCES kategori(id)
ON DELETE SET NULL
ON UPDATE CASCADE;

-- Update existing data to set default kategori_id (optional)
UPDATE produk SET kategori_id = 1 WHERE kategori_id IS NULL;
