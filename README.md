# SOLID Principle
The SOLID principle are a set of five designs guidelines that help developers create more maintainable, flexible, and scalable software.

Here's an overview of SOLID principles:
### 1. Single Responsibility Principle (SRP)
A class should have one, and only one, reason to change.
##### Example Case:
- Before Implementation
```go
type OrderServiceNotSOP struct {
    Orders []entity.Order
}

func (o *OrderServiceNotSOP) PlaceOrder(order entity.Order) {
    o.Orders = append(o.Orders, order)
}

func (o *OrderServiceNotSOP) GenerateInvoice(order entity.Order) {
    fmt.Printf("Generate Invoice For Order ID: %s, User ID: %d and Total Amount: %f", order.OrderID, order.UserID, order.TotalAmount)
}
```
OrderServiceNotSOP has two responsibilities: handling orders and generating invoices.

- After Implementation
```go
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
```
Now, the OrderServiceSOP is solely responsible for managing orders, while the InvoiceGeneratorSOP takes care of generating invoices.

### 2. Open/Closed Principle (OCP)
Software entities should be open for extension but closed for modification.
##### Example Case:
// TODO

### 3. Liskov Substitution Principle (LSP)
Subtypes must be substitutable for their base types without altering the correctness of the program.
##### Example Case:
// TODO

### 4. Interface Segregation Principle (ISP)
No client should be forced to depend on interfaces it does not use.
##### Example Case:
// TODO

### 5. Dependency Inversion Principle (DIP)
High-level modules should not depend on low-level modules. Both should depend on abstractions.
##### Example Case:
// TODO