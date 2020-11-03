package invoiceitem

import "time"

//Model of invoice item
type Model struct {
	Id uint
	InvoiceHeaderID uint
	ProductId uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
