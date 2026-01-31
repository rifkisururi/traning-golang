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

	// Database configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// GetDBConnectionString mengembalikan connection string untuk PostgreSQL
func (c *Config) GetDBConnectionString() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=" + c.DBSSLMode
}

// LoadConfig memuat konfigurasi dari environment variables atau file .env
// Prioritas: environment variables > file .env > default values
func LoadConfig() *Config {
	v := viper.New()

	// Set default values
	v.SetDefault("HOST", "localhost")
	v.SetDefault("PORT", "8080")

	// Database default values
	v.SetDefault("DB_HOST", "localhost")
	v.SetDefault("DB_PORT", "5432")
	v.SetDefault("DB_USER", "postgres")
	v.SetDefault("DB_PASSWORD", "")
	v.SetDefault("DB_NAME", "kasir_db")
	v.SetDefault("DB_SSLMODE", "disable")

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
	v.BindEnv("DB_HOST")
	v.BindEnv("DB_PORT")
	v.BindEnv("DB_USER")
	v.BindEnv("DB_PASSWORD")
	v.BindEnv("DB_NAME")
	v.BindEnv("DB_SSLMODE")

	// Membaca konfigurasi
	config := &Config{
		Host:       v.GetString("HOST"),
		Port:       v.GetString("PORT"),
		DBHost:     v.GetString("DB_HOST"),
		DBPort:     v.GetString("DB_PORT"),
		DBUser:     v.GetString("DB_USER"),
		DBPassword: v.GetString("DB_PASSWORD"),
		DBName:     v.GetString("DB_NAME"),
		DBSSLMode:  v.GetString("DB_SSLMODE"),
	}

	log.Printf("[config] Konfigurasi dimuat - Host: %s, Port: %s", config.Host, config.Port)
	log.Printf("[config] Database - Host: %s, Port: %s, DB: %s", config.DBHost, config.DBPort, config.DBName)

	return config
}
