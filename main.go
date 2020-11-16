package main

import (
	"github.com/cmorales95/golang_db/pkg/invoice"
	"github.com/cmorales95/golang_db/pkg/invoiceheader"
	"github.com/cmorales95/golang_db/pkg/invoiceitem"
	"github.com/cmorales95/golang_db/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alexys",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatal("invoice.create: %v", err)
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

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceHeader := invoiceheader.NewService(storageHeader)
	serviceItems := invoiceitem.NewService(storageItems)

	if err := serviceHeader.Migrate(); err != nil {
		log.Fatalf("%v", err)
	}
	if err := serviceItems.Migrate(); err != nil {
		log.Fatalf("%v", err)
	}


storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.delete: %v", err)
	}
*/