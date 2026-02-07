-- Drop index transaction_details terlebih dahulu.
DROP INDEX IF EXISTS idx_transaction_details_transaction_id;

-- Drop tabel transaction_details.
DROP TABLE IF EXISTS transaction_details;

-- Drop tabel transactions.
DROP TABLE IF EXISTS transactions;
