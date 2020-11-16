package main

import (
	"github.com/cmorales95/golang_db/pkg/invoice"
	"github.com/cmorales95/golang_db/pkg/invoiceheader"
	"github.com/cmorales95/golang_db/pkg/invoiceitem"
	"github.com/cmorales95/golang_db/storage"
	"log"
)

func main() {
	storage.NewMySQLDB()
	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Cristian Morales",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 3},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatal("invoice.create: %v", err)
	}

}

/*

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceHeader := invoiceheader.NewService(storageHeader)
	serviceItems := invoiceitem.NewService(storageItems)

	if err := serviceHeader.Migrate(); err != nil {
		log.Fatalf("%v", err)
	}
	if err := serviceItems.Migrate(); err != nil {
		log.Fatalf("%v", err)
	}


storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.delete: %v", err)
	}


storage.NewPostgresDB()

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
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
*/
