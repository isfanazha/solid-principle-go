package singleresponsibleprinciple

import (
	"fmt"

	"github.com/isfanazha/solid-principle-go/entity"
)

// PROBLEM
// OrderServiceNotSOP has two responsibilities: handling orders and generating invoices.

type OrderServiceNotSOP struct {
	Orders []entity.Order
}

func (o *OrderServiceNotSOP) PlaceOrder(order entity.Order) {
	o.Orders = append(o.Orders, order)
}

func (o *OrderServiceNotSOP) GenerateInvoice(order entity.Order) {
	fmt.Printf("Generate Invoice For Order ID: %s, User ID: %d and Total Amount: %f", order.OrderID, order.UserID, order.TotalAmount)
}
