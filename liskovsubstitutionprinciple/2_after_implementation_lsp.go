package liskovsubstitutionprinciple

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

// RegularProductLSP is excluded from the discountable behavior by not implementing the Discountable interface.
// This design clearly separates products that can have discounts from those that cannot, reducing confusion
// and making the code easier to maintain and extend.

func ApplyDiscounts(products []Discountable) {
	for _, product := range products {
		product.ApplyDiscount()
	}
}
