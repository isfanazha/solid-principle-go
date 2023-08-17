package singleresponsibleprinciple

import (
	"fmt"

	"github.com/isfanazha/solid-principle-go/entity"
)

// SOLUTION
// OrderServiceSOP is solely responsible for managing orders, while the InvoiceGeneratorSOP takes care of generating invoices.

type OrderServiceSOP struct {
	Orders []entity.Order
}

func (o *OrderServiceSOP) PlaceOrder(order entity.Order) {
	o.Orders = append(o.Orders, order)
}

type InvoiceGeneratorSOP struct{}

func (i *InvoiceGeneratorSOP) GenerateInvoice(order entity.Order) {
	fmt.Printf("Generate Invoice For Order ID: %s, User ID: %d and Total Amount: %f", order.OrderID, order.UserID, order.TotalAmount)
}
