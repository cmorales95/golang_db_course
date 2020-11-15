package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/product"
	"github.com/cmorales95/golang_db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)


	m, err := serviceProduct.GetById(3)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There is no product with the id")
	case err != nil :
			log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(m)
	}
}


/*
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("Error")
	}
*/