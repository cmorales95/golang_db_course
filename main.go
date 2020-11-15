package main

import (
	"github.com/cmorales95/golang_db/pkg/product"
	"github.com/cmorales95/golang_db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.delete: %v", err)
	}
}


/*
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("Error")
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There is no product with the id")
	case err != nil :
			log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(m)
	}

	m := &product.Model{
		ID:           5,
		Name:         "Curso de Go",
		Price: 90,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.update: %v", err)
	}
*/