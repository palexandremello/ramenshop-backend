package postegresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitializeDatabase() *sql.DB {

	host := "172.18.0.1"
	port := 5432
	user := "ramenshop_backend"
	password := "ramenshop_password"
	dbname := "ramenshop_db"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	return db
}
