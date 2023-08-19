package openclosedprinciple

import "github.com/isfanazha/solid-principle-go/entity"

// PROBLEM
// if you want to add more discount rules, you'll have to modify the ApplyDiscount function, violating the OCP.

func ApplyDiscount(product entity.Product, isVIP bool) float64 {
	if isVIP {
		return product.Price * 0.20 // 20% discount for VIP
	}

	return product.Price * 0.10 // 10% discount otherwise
}
