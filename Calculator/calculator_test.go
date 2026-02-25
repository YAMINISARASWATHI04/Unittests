package calculator

import (
	// "calculator/Calculator"
	"testing"
	"github.com/stretchr/testify/assert"
	// "calculator/Calculator/database"
)

type DiscountRepositoryMock struct{

}

func (drm *DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscountCalculator(t *testing.T){
	type testCase struct{
		name string
		minimumPurchaseAmount int
		
		purchaseAmount int
		expectedAmount int
	}
	tests := []testCase{
		{name : "a" ,minimumPurchaseAmount: 100,purchaseAmount: 150,expectedAmount: 130},
		{name : "b" ,minimumPurchaseAmount: 100,purchaseAmount: 200,expectedAmount: 160},
		{name : "c" ,minimumPurchaseAmount: 100,purchaseAmount: 350,expectedAmount: 290},
		{name : "d" ,minimumPurchaseAmount: 100,purchaseAmount: 50,expectedAmount: 50},
		// {name : "e" ,minimumPurchaseAmount: 0, purchaseAmount: 50,expectedAmount: 50},
	}

	for _,tc :=range tests{
		t.Run(tc.name,func (t *testing.T){

			discrepomock :=DiscountRepositoryMock{}
			calculator ,err :=NewDiscountCalculator(tc.minimumPurchaseAmount,&discrepomock)

			if err!=nil{
				// fail+log
				//t.Errorf("couldnt instantiate calculator")
				//FailNow- It will not execute further actions
				t.Fatal("couldnt instantiate calculator")
			}
			amount :=calculator.Calculate(tc.purchaseAmount)

			// if amount!=tc.expectedAmount{
			// 	t.Errorf("Expected amount %v, got %v",tc.expectedAmount,amount)
			// }
			assert.Equal(t,tc.expectedAmount,amount)
		})
		
	}
}
