package openclosedprinciple

import "github.com/isfanazha/solid-principle-go/entity"

// SOLUTION
// Define an interface for discount rules and then implement specific rules as separate types.

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
