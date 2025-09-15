package handlers

import (
	"fmt"
	"os"

	"golang-toko/migrations"
	"golang-toko/seeders"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// GetDB mengembalikan instance GORM DB.
func GetDB() *gorm.DB {
	return db
}

func init() {
	var err error

	// Ambil kredensial database dari variabel lingkungan.
	// Gunakan nilai default jika variabel lingkungan tidak ditemukan.
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = "Zaghi08@"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "ecomm"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	// Buat string koneksi menggunakan variabel lingkungan
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Buka koneksi ke MySQL. Jika gagal, program akan panik di sini.
	db, err = gorm.Open("mysql", dbURI)
	if err != nil {
		panic("Gagal terhubung ke database MySQL: " + err.Error())
	}

	// Lakukan ping untuk memastikan koneksi valid.
	if pingErr := db.DB().Ping(); pingErr != nil {
		panic("Gagal ping database: " + pingErr.Error())
	}

	// Jalankan migrasi dan seeder.
	fmt.Println("[Handlers] Koneksi database berhasil. Menjalankan migrasi dan seeder...")
	migrations.Migrate(db)
	seeders.Seed(db)
}

// package handlers

// import (
// 	"fmt"
// 	"golang-toko/migrations"
// 	"golang-toko/seeders"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// var db *gorm.DB

// func init() {
// 	var err error

// 	// Try primary MySQL database first.
// 	db, err = gorm.Open("mysql", "root:Zaghi08@@tcp(127.0.0.1:3306)/ecomm?charset=utf8mb4&parseTime=True&loc=Local")

// 	// Check connectivity – db.Ping reveals whether the database is actually reachable.
// 	if err == nil {
// 		if pingErr := db.DB().Ping(); pingErr != nil {
// 			err = pingErr
// 		}
// 	}

// 	if err != nil {
// 		// Fallback to an in-memory SQLite database for local/testing environments.
// 		db, err = gorm.Open("sqlite3", ":memory:")
// 		if err != nil {
// 			panic("failed to connect database: " + err.Error())
// 		}
// 		// Keep a single connection so the in-memory database persists for the life of the process.
// 		db.DB().SetMaxOpenConns(1)
// 	}

// 	// Run migrations and seed data so that tests have the required tables & records.
// 	fmt.Println("[handlers] running migrations and seeding test data")
// 	migrations.Migrate(db)
// 	seeders.Seed(db)
// }
