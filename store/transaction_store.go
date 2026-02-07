package store

import (
	"database/sql"
	"fmt"
	"log"

	"kasir-api/database"
	"kasir-api/models"
)

// CreateTransaction membuat transaksi baru beserta detailnya dalam satu database transaction.
func CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	// Mulai database transaction.
	tx, err := database.DB.Begin()
	if err != nil {
		log.Printf("[transaction-store] Error begin transaction: %v", err)
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	// Proses setiap item: validasi produk, hitung subtotal, kurangi stok.
	for _, item := range items {
		var productPrice, stock int
		var productName string

		// Ambil data produk dan cek stok.
		err := tx.QueryRow("SELECT nama, harga, stok FROM produk WHERE id = $1", item.ProductID).
			Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			log.Printf("[transaction-store] Product not found id=%d", item.ProductID)
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			log.Printf("[transaction-store] Error get product: %v", err)
			return nil, err
		}

		// Validasi stok cukup.
		if stock < item.Quantity {
			log.Printf("[transaction-store] Insufficient stock product_id=%d requested=%d available=%d",
				item.ProductID, item.Quantity, stock)
			return nil, fmt.Errorf("insufficient stock for product %s (requested: %d, available: %d)",
				productName, item.Quantity, stock)
		}

		// Hitung subtotal.
		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// Kurangi stok produk.
		_, err = tx.Exec("UPDATE produk SET stok = stok - $1 WHERE id = $2",
			item.Quantity, item.ProductID)
		if err != nil {
			log.Printf("[transaction-store] Error update stock: %v", err)
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Insert transaction record dan dapatkan ID.
	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id",
		totalAmount).Scan(&transactionID)
	if err != nil {
		log.Printf("[transaction-store] Error insert transaction: %v", err)
		return nil, err
	}

	// Insert transaction details dan dapatkan ID masing-masing detail.
	for i := range details {
		details[i].TransactionID = transactionID
		err := tx.QueryRow(
			"INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal,
		).Scan(&details[i].ID)
		if err != nil {
			log.Printf("[transaction-store] Error insert transaction detail: %v", err)
			return nil, err
		}
	}

	// Commit transaction.
	if err := tx.Commit(); err != nil {
		log.Printf("[transaction-store] Error commit transaction: %v", err)
		return nil, err
	}

	log.Printf("[transaction-store] Transaction created id=%d total=%d items=%d",
		transactionID, totalAmount, len(details))

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

// GetTransactionByID mengembalikan satu transaksi berdasarkan ID beserta detailnya.
func GetTransactionByID(id int) (*models.Transaction, error) {
	var transaction models.Transaction

	// Ambil data transaksi.
	err := database.DB.QueryRow("SELECT id, total_amount, created_at FROM transactions WHERE id = $1", id).
		Scan(&transaction.ID, &transaction.TotalAmount, &transaction.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("transaction id %d not found", id)
	}
	if err != nil {
		log.Printf("[transaction-store] Error get transaction: %v", err)
		return nil, err
	}

	// Ambil detail transaksi dengan join ke produk untuk nama produk.
	rows, err := database.DB.Query(`
		SELECT td.id, td.transaction_id, td.product_id, p.nama, td.quantity, td.subtotal
		FROM transaction_details td
		JOIN produk p ON td.product_id = p.id
		WHERE td.transaction_id = $1
		ORDER BY td.id
	`, id)
	if err != nil {
		log.Printf("[transaction-store] Error get transaction details: %v", err)
		return nil, err
	}
	defer rows.Close()

	var details []models.TransactionDetail
	for rows.Next() {
		var d models.TransactionDetail
		if err := rows.Scan(&d.ID, &d.TransactionID, &d.ProductID, &d.ProductName, &d.Quantity, &d.Subtotal); err != nil {
			log.Printf("[transaction-store] Error scanning detail row: %v", err)
			continue
		}
		details = append(details, d)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[transaction-store] Error iterating detail rows: %v", err)
		return nil, err
	}

	transaction.Details = details
	return &transaction, nil
}

// GetAllTransactions mengembalikan semua transaksi (tanpa detail untuk performa).
func GetAllTransactions() ([]models.Transaction, error) {
	rows, err := database.DB.Query("SELECT id, total_amount, created_at FROM transactions ORDER BY id DESC")
	if err != nil {
		log.Printf("[transaction-store] Error get all transactions: %v", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.TotalAmount, &t.CreatedAt); err != nil {
			log.Printf("[transaction-store] Error scanning transaction row: %v", err)
			continue
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[transaction-store] Error iterating transaction rows: %v", err)
		return nil, err
	}

	return transactions, nil
}
