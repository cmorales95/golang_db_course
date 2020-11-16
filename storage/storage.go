package storage

import (
	"database/sql"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/product"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

var (
	db *sql.DB
	once sync.Once
)

//Driver of storage
type Driver string

const (
	Mysql    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

//New create a connection with DB
func New(d Driver) {
	switch d {
	case Mysql:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newPostgresDB singleton
func newPostgresDB() {
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

// newMySQLDB singleton
func newMySQLDB() {
	once.Do(func(){
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("Connection with DB is not available: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Connection with Ping DB: %v", err)
		}
		fmt.Println("Connected to MySQL!")
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

//DAOProduct factory of product.Storage
func DAOProduct(d Driver) (product.Storage, error) {
	switch d {
	case Postgres:
		return newPsqlProduct(db), nil
	case Mysql:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("driver not implemented")
	}
}