package main

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config menyimpan konfigurasi aplikasi
type Config struct {
	Host string
	Port string
}

// LoadConfig memuat konfigurasi dari environment variables atau file .env
// Prioritas: environment variables > file .env > default values
func LoadConfig() *Config {
	v := viper.New()

	// Set default values
	v.SetDefault("HOST", "localhost")
	v.SetDefault("PORT", "8080")

	// Konfigurasi untuk membaca file .env
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("..")

	// Cek file .env ada atau tidak
	if _, err := os.Stat(".env"); err == nil {
		// File .env ditemukan, baca file tersebut
		if err := v.ReadInConfig(); err != nil {
			log.Printf("[config] Warning: gagal membaca file .env: %v", err)
		} else {
			log.Printf("[config] File .env berhasil dibaca")
		}
	}

	// AutomaticEnv akan membuat viper mengecek environment variables
	// Jika env variable ditemukan, akan override nilai dari .env atau default
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Bind environment variables
	v.BindEnv("HOST")
	v.BindEnv("PORT")

	// Membaca konfigurasi
	config := &Config{
		Host: v.GetString("HOST"),
		Port: v.GetString("PORT"),
	}

	log.Printf("[config] Konfigurasi dimuat - Host: %s, Port: %s", config.Host, config.Port)

	return config
}
