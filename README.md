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
```go
type ProductNotLSP interface {
    ApplyDiscount()
}

type DiscountedProductNotLSP struct {
    Name  string
    Price float64
}

func (dp *DiscountedProductNotLSP) ApplyDiscount() {
    dp.Price *= 0.90 // 10% discount
}

type RegularProductNotLSP struct {
    Name  string
    Price float64
}

func (rp *RegularProductNotLSP) ApplyDiscount() {
    // No discount for regular products
}

func ApplyDiscountsNotLSP(products []ProductNotLSP) {
    for _, product := range products {
        product.ApplyDiscount()
    }
}
```
The issue here is that this design doesn't make it clear which products are eligible for a discount, leading to confusion and potential errors. If a new category is added, the discount logic might need to change in many places.

- After Implementation
```go
type Discountable interface {
    ApplyDiscount()
}

type ProductLSP struct {
    Name  string
    Price float64
}

type DiscountedProductLSP struct {
    ProductLSP
}

func (dp *DiscountedProductLSP) ApplyDiscount() {
    dp.Price *= 0.90 // 10% discount
}

func ApplyDiscounts(products []Discountable) {
    for _, product := range products {
        product.ApplyDiscount()
    }
}
```
`RegularProductLSP` is excluded from the discountable behavior by not implementing the `Discountable` interface.  This design clearly separates products that can have discounts from those that cannot, reducing confusion and making the code easier to maintain and extend.
Let's breakdown the code based on the definition:
- Base types (Interface): `Discountable`interface.
- Struct (Possible Subtype): `ProductLSP` struct.
- Embedded Struct (Subtype): The `DiscountedProductLSP` struct embeds `ProductLSP`. This effectively means `DiscountedProductLSP` inherits properties from `ProductLSP` and then adds its own functionality, specifically the `ApplyDiscount()` method, making it adhere to the `Discountable` interface. Hence, `DiscountedProductLSP` is a subtype of the `Discountable` interface.

##### Liskov Substitution Principle in Action: 
Given the LSP, you should be able to replace an instance of the `Discountable` interface (base type) with an instance of `DiscountedProductLSP` (subtype) without affecting the program's correctness.
```
func ApplyDiscounts(products []Discountable) {
    for _, product := range products {
        product.ApplyDiscount()
    }
}
```
You can pass a slice of `Discountable` items, and as long as every item in that slice adheres to the `Discountable` interface (i.e., it has an `ApplyDiscount()` method), the function will work correctly. Here, `DiscountedProductLSP` can be used as a substitutable type for `Discountable` because it implements the required method.

This means that, if in the future you introduce another product type with a different discount strategy (say, a `ClearanceProduct`), as long as it implements the Discountable interface, you can pass it to the `ApplyDiscounts` function without any issues. This is the essence of the Liskov Substitution Principle.


### 4. Interface Segregation Principle (ISP)
No client should be forced to depend on interfaces it does not use.
##### Example Case:
// TODO

### 5. Dependency Inversion Principle (DIP)
High-level modules should not depend on low-level modules. Both should depend on abstractions.
##### Example Case:
// TODO