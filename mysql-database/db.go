package main

import (
	"database/sql"
	"encoding/hex"
	"log"
	"time"

	"github.com/codahale/newplex/mhf"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@(127.0.0.1:3306)/test?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	createTable(db)

	insertUser(db)

	selectByUsername(db)

	deleteByUsername(db)
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id int AUTO_INCREMENT,
			username text NOT NULL,
			password text NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);
	`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatalf(err.Error())
	}

}

func insertUser(db *sql.DB) {
	username := "renato"
	password := createPassword()
	created_at := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, created_at)

	if err != nil {
		log.Fatalf(err.Error())
	}

	userId, err := result.LastInsertId()

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("id: %d", userId)
}

func createPassword() string {
	password := "test123"
	domain := "test"
	cost := uint8(10)
	salt := "test-gowebexamples"
	n := 32

	hash := mhf.Hash(domain, cost, []byte(salt), []byte(password), nil, n)

	return hex.EncodeToString(hash)
}

func selectByUsername(db *sql.DB) {
	var (
		username string
		password string
	)

	query := `SELECT username, password FROM users WHERE username = ?`

	err := db.QueryRow(query, "renato").Scan(&username, &password)

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("username found %s", username)
}

func deleteByUsername(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users WHERE username = ?`, "renato")

	if err != nil {
		log.Fatalf(err.Error())
	}
}
