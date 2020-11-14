package main

import (
	"fmt"
	"github.com/cmorales95/golang_db/pkg/product"
	"github.com/cmorales95/golang_db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Curso de POO con Go",
		Price:        70,
		Observations: "on fine",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	fmt.Printf("%+v\n", m)
}
