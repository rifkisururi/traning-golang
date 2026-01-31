package store

import (
	"database/sql"
	"log"

	"kasir-api/database"
	"kasir-api/models"
)

// GetAllKategori mengembalikan semua kategori
func GetAllKategori() []models.Kategori {
	rows, err := database.DB.Query("SELECT id, nama, deskripsi FROM kategori ORDER BY id")
	if err != nil {
		log.Printf("[kategori-store] Error GetAllKategori: %v", err)
		return []models.Kategori{}
	}
	defer rows.Close()

	var kategori []models.Kategori
	for rows.Next() {
		var k models.Kategori
		if err := rows.Scan(&k.ID, &k.Nama, &k.Deskripsi); err != nil {
			log.Printf("[kategori-store] Error scanning row: %v", err)
			continue
		}
		kategori = append(kategori, k)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[kategori-store] Error iterating rows: %v", err)
	}

	return kategori
}

// GetKategoriByID mencari kategori berdasarkan ID
func GetKategoriByID(id int) (models.Kategori, bool) {
	var k models.Kategori
	err := database.DB.QueryRow("SELECT id, nama, deskripsi FROM kategori WHERE id = $1", id).
		Scan(&k.ID, &k.Nama, &k.Deskripsi)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Kategori{}, false
		}
		log.Printf("[kategori-store] Error GetKategoriByID: %v", err)
		return models.Kategori{}, false
	}

	return k, true
}

// AddKategori menambahkan kategori baru dan mengembalikan kategori dengan ID
func AddKategori(k models.Kategori) (models.Kategori, error) {
	err := database.DB.QueryRow(
		"INSERT INTO kategori (nama, deskripsi) VALUES ($1, $2) RETURNING id",
		k.Nama, k.Deskripsi,
	).Scan(&k.ID)

	if err != nil {
		log.Printf("[kategori-store] Error AddKategori: %v", err)
		return models.Kategori{}, err
	}

	return k, nil
}

// UpdateKategori mengupdate kategori yang sudah ada
func UpdateKategori(id int, updated models.Kategori) bool {
	result, err := database.DB.Exec(
		"UPDATE kategori SET nama = $1, deskripsi = $2 WHERE id = $3",
		updated.Nama, updated.Deskripsi, id,
	)

	if err != nil {
		log.Printf("[kategori-store] Error UpdateKategori: %v", err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}

// DeleteKategori menghapus kategori berdasarkan ID
func DeleteKategori(id int) bool {
	result, err := database.DB.Exec("DELETE FROM kategori WHERE id = $1", id)

	if err != nil {
		log.Printf("[kategori-store] Error DeleteKategori: %v", err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}
