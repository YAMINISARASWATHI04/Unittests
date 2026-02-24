package calculator

type DiscountCalculator struct{
	minpurchase int
	discountamount int
}

func NewDiscountCalculator(minpurchase int,discountamount int) *DiscountCalculator{
	return &DiscountCalculator{
		minpurchase:minpurchase,
		discountamount: discountamount,

	}
}

func (c *DiscountCalculator) Calculate(purchaseamount int) int {
	if purchaseamount>c.minpurchase {
		multiplier :=purchaseamount/c.minpurchase
		return purchaseamount-c.discountamount*multiplier
	}

	
	return purchaseamount
}
