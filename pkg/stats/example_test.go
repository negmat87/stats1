package stats

import (
	"github.com/negmat87/bank/pkg/types"
	"fmt"
)
func ExampleAvg() {
	
	paymtnts := []types.Payment{
		{
			ID:  1,
			Amount: 5_00_00,
		},
		{
			ID:  2,
			Amount: 6_00_00,
		},
		{
			ID:  3,
			Amount: 5_000_00,
		},
		{
			ID:  4,
			Amount: 3_00_00,
		},
	}

	sum :=Avg(paymtnts)
	fmt.Println(sum)
	
	//Output:
	//160000
}

func ExampleTotalInCategory() {
	
	paymtnts := []types.Payment{
		{
			ID:  1,
			Amount: 5_00_00,
			Category: types.STUDY,
		},
		{
			ID:  2,
			Amount: 6_00_00,
			Category: types.STUDY,
		},
		{
			ID:  3,
			Amount: 5_000_00,
			Category: types.FOOD,
		},
		{
			ID:  4,
			Amount: 3_00_00,
			Category: types.CLOTHES,
		},
	}

	category := "STUDY"	
	sumCategory :=TotalInCategory(paymtnts, types.Category(category))
	
	fmt.Println(sumCategory)
	
	//Output:
	//110000
}