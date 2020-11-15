package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

var (
	db *sql.DB
	once sync.Once
)

// NewPostgresDB singleton
func NewPostgresDB() {
	once.Do(func(){
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Connection with DB is not available: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Connection with Ping DB: %v", err)
		}
		fmt.Println("Connected to Postgres!")
	})
}

// Pool return a unique instance of DB
func Pool() *sql.DB {
	return db
}

//stringToNull control of null value for string
func stringToNull(s string) sql.NullString {
	null := sql.NullString{
		String: s,
	}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{
		Time: t,
	}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}