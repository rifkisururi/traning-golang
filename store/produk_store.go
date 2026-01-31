package store

import (
	"database/sql"
	"log"

	"kasir-api/database"
	"kasir-api/models"
)

// GetAll mengembalikan semua data produk.
func GetAll() []models.Produk {
	rows, err := database.DB.Query("SELECT id, nama, harga, stok FROM produk ORDER BY id")
	if err != nil {
		log.Printf("[produk-store] Error GetAll: %v", err)
		return []models.Produk{}
	}
	defer rows.Close()

	var produk []models.Produk
	for rows.Next() {
		var p models.Produk
		if err := rows.Scan(&p.ID, &p.Nama, &p.Harga, &p.Stok); err != nil {
			log.Printf("[produk-store] Error scanning row: %v", err)
			continue
		}
		produk = append(produk, p)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[produk-store] Error iterating rows: %v", err)
	}

	return produk
}

// GetByID mengembalikan satu produk berdasarkan ID.
func GetByID(id int) (models.Produk, bool) {
	var p models.Produk
	err := database.DB.QueryRow("SELECT id, nama, harga, stok FROM produk WHERE id = $1", id).
		Scan(&p.ID, &p.Nama, &p.Harga, &p.Stok)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Produk{}, false
		}
		log.Printf("[produk-store] Error GetByID: %v", err)
		return models.Produk{}, false
	}

	return p, true
}

// Add menambahkan produk baru ke penyimpanan dan mengembalikan produk dengan ID
func Add(p models.Produk) (models.Produk, error) {
	err := database.DB.QueryRow(
		"INSERT INTO produk (nama, harga, stok) VALUES ($1, $2, $3) RETURNING id",
		p.Nama, p.Harga, p.Stok,
	).Scan(&p.ID)

	if err != nil {
		log.Printf("[produk-store] Error Add: %v", err)
		return models.Produk{}, err
	}

	return p, nil
}

// Update mengganti data produk berdasarkan ID.
func Update(id int, p models.Produk) (models.Produk, bool) {
	result, err := database.DB.Exec(
		"UPDATE produk SET nama = $1, harga = $2, stok = $3 WHERE id = $4",
		p.Nama, p.Harga, p.Stok, id,
	)

	if err != nil {
		log.Printf("[produk-store] Error Update: %v", err)
		return models.Produk{}, false
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		p.ID = id
		return p, true
	}

	return models.Produk{}, false
}

// Delete menghapus produk berdasarkan ID.
func Delete(id int) bool {
	result, err := database.DB.Exec("DELETE FROM produk WHERE id = $1", id)

	if err != nil {
		log.Printf("[produk-store] Error Delete: %v", err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}
