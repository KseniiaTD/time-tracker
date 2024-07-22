package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func (db *Database) Disconnect() {
	db.db.Close()
}

func (db *Database) DB() *sql.DB {
	return db.db
}

func Connect() (Database, error) {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PWD")
	dbname := os.Getenv("PG_DBNAME")
	conn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var db Database
	database, err := sql.Open("postgres", conn)

	if err != nil {
		return Database{}, err
	}

	db.db = database

	err = db.db.Ping()
	if err != nil {
		return Database{}, err
	}

	return db, nil
}
