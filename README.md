

# Principal Execution

#### Migrate
```
        storage.NewPostgresDB()
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}
```

#### Create a Product

```
    storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:         "Curso de test con Go",
		Price:        90,
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.create: %v", err)
	}
		fmt.Printf("%+v\n", m)
``` 

#### Get All
```
    storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)


	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("Error")
	}
	fmt.Println(ms)
```

#### Get by ID
```
	storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetById(1)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("There is no product with the id")
	case err != nil :
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(m)
	}
	fmt.Println(m)
```

 #### Update 
 
 ```
    storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:    1,
		Name:  "Curso de Go",
		Price: 85,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.update: %v", err)
	}
```

#### Delete

```
        storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.update: %v", err)
	}
```

#### Create Invoce with Details
```
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
```