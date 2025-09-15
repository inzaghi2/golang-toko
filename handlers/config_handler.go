package handlers

import (
	"fmt"
	"golang-toko/migrations"
	"golang-toko/seeders"
	"os" // Tambahkan package "os" untuk membaca variabel lingkungan

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error

	// Baca variabel lingkungan untuk kredensial database MySQL
	dbHost := os.Getenv("127.0.0.1")
	dbUser := os.Getenv("root")
	dbPass := os.Getenv("Zaghi08@")
	dbName := os.Getenv("ecomm")
	dbPort := os.Getenv("3306")

	// Pastikan semua variabel lingkungan ada
	if dbHost == "" || dbUser == "" || dbPass == "" || dbName == "" {
		// Jika ada variabel yang kosong, berikan pesan error yang jelas
		// dan hentikan program. Ini lebih baik daripada panic yang tidak jelas.
		fmt.Println("Error: Variabel lingkungan database tidak lengkap. Pastikan DB_HOST, DB_USER, DB_PASS, dan DB_NAME sudah disetel.")
		panic("Database environment variables are not set.")
	}

	// Buat string koneksi MySQL menggunakan variabel lingkungan
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Buka koneksi ke database MySQL
	db, err = gorm.Open("mysql", dbURI)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Ping database untuk memastikan koneksi berhasil
	if pingErr := db.DB().Ping(); pingErr != nil {
		panic("failed to ping database: " + pingErr.Error())
	}

	// Run migrations and seed data
	fmt.Println("[handlers] running migrations and seeding data")
	migrations.Migrate(db)
	seeders.Seed(db)
}
