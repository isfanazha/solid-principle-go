package liskovsubstitutionprinciple

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
	// The issue here is that this design doesn't make it clear which products are eligible for a discount,
	// leading to confusion and potential errors. If a new category is added, the discount logic might need
	// to change in many places.
}

func ApplyDiscountsNotLSP(products []ProductNotLSP) {
	for _, product := range products {
		product.ApplyDiscount()
	}
}
