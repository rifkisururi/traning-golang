-- Rollback: Remove kategori_id column and foreign key from produk table

-- Drop foreign key constraint
ALTER TABLE produk DROP CONSTRAINT IF EXISTS fk_produk_kategori;

-- Drop kategori_id column
ALTER TABLE produk DROP COLUMN IF EXISTS kategori_id;
