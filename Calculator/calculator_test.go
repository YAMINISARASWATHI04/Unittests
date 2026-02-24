package calculator

import (
	// "calculator/Calculator"
	
	"testing"
)

func TestDiscountCalculator(t *testing.T){
	type testCase struct{
		name string
		minimumPurchaseAmount int
		discount int
		purchaseAmount int
		expectedAmount int
	}
	tests := []testCase{
		{name : "a" ,minimumPurchaseAmount: 100,discount: 20,purchaseAmount: 150,expectedAmount: 130},
		{name : "b" ,minimumPurchaseAmount: 100,discount: 10,purchaseAmount: 200,expectedAmount: 160},
		{name : "c" ,minimumPurchaseAmount: 100,discount: 20,purchaseAmount: 350,expectedAmount: 290},
		{name : "d" ,minimumPurchaseAmount: 100,discount: 20,purchaseAmount: 50,expectedAmount: 50},
	}

	for _,tc :=range tests{
		t.Run(tc.name,func (t *testing.T){
			calculator:=NewDiscountCalculator(tc.minimumPurchaseAmount,tc.discount)
			amount :=calculator.Calculate(tc.purchaseAmount)

			if amount!=tc.expectedAmount{
				t.Errorf("Expected amount %v, got %v",tc.expectedAmount,amount)
			}
		})
		
	}
}
