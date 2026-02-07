-- Membuat tabel transactions untuk menyimpan data transaksi kasir.
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel transaction_details untuk menyimpan detail item dalam setiap transaksi.
CREATE TABLE IF NOT EXISTS transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES produk(id),
    quantity INT NOT NULL,
    subtotal INT NOT NULL
);

-- Membuat index untuk performa query transaction_details berdasarkan transaction_id.
CREATE INDEX IF NOT EXISTS idx_transaction_details_transaction_id ON transaction_details(transaction_id);
