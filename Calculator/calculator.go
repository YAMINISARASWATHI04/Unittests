package calculator
import (
	"errors"
	"calculator/Calculator/database"
)

type DiscountCalculator struct{
	minpurchase int
	discountrepo database.Repository
}

func NewDiscountCalculator(minpurchase int,discountrepo database.Repository) (*DiscountCalculator,error){
	if minpurchase==0{
		return &DiscountCalculator{} ,errors.New("Purchase amount cannot be zero")
	}
	return &DiscountCalculator{
		minpurchase:minpurchase,
		discountrepo: discountrepo,

	},nil
}

func (c *DiscountCalculator) Calculate(purchaseamount int) int {
	if purchaseamount>c.minpurchase {
		multiplier :=purchaseamount/c.minpurchase
		discount :=c.discountrepo.FindCurrentDiscount()
		return purchaseamount-(discount*multiplier)
	}

	
	return purchaseamount
}
