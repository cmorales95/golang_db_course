package main

import (
	"fmt"
	"github.com/cmorales95/golang_db/pkg/product"
	"github.com/cmorales95/golang_db/storage"
	"log"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("daoproduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)
	m, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("daoproduct: %v", err)
	}
	fmt.Println(m)
}
