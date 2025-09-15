package handlers

import (
	"fmt"
	"golang-toko/migrations"
	"golang-toko/seeders"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error

	// Try primary MySQL database first.
	db, err = gorm.Open("mysql", "root:AzBlfEsMNczEdwYEOozgKZwKPajpqMuK@crossover.proxy.rlwy.net:26161/railway")

	// Check connectivity – db.Ping reveals whether the database is actually reachable.
	if err == nil {
		if pingErr := db.DB().Ping(); pingErr != nil {
			err = pingErr
		}
	}

	if err != nil {
		// Fallback to an in-memory SQLite database for local/testing environments.
		db, err = gorm.Open("sqlite3", ":memory:")
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
		// Keep a single connection so the in-memory database persists for the life of the process.
		db.DB().SetMaxOpenConns(1)
	}

	// Run migrations and seed data so that tests have the required tables & records.
	fmt.Println("[handlers] running migrations and seeding test data")
	migrations.Migrate(db)
	seeders.Seed(db)
}
