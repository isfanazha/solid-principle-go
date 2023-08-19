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
- Before Implementation
```go
func ApplyDiscount(product entity.Product, isVIP bool) float64 {
    if isVIP {
        return product.Price * 0.20 // 20% discount for VIP
    }

    return product.Price * 0.10 // 10% discount otherwise
}

```
Imagine you have a simple discount system that applies a 10% discount to all products. Later, you want to add a special 20% discount for VIP customers.

- After Implementation
```go
type DiscountRule interface {
    ApplyDiscount(product entity.Product) float64
}

type StandardDiscount struct{}

func (sd StandardDiscount) ApplyDiscount(product entity.Product) float64 {
    return product.Price * 0.10 // 10% discount
}

type VIPDiscount struct{}

func (vd VIPDiscount) ApplyDiscount(product entity.Product) float64 {
    return product.Price * 0.20 // 20% discount for VIP
}

func CalculatePrice(product entity.Product, rule DiscountRule) float64 {
    return product.Price - rule.ApplyDiscount(product)
}
```
Now, when you want to add a new discount rule, you simply create a new type that implements the `DiscountRule` interface. The existing code doesn't have to change, adhering to OCP.

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