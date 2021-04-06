package stats

import (
	"github.com/negmat87/bank/v2/pkg/types"
)

func Avg(paymtnts []types.Payment) int64 {
	sum := int64(0)
	avg := len(paymtnts)
	for _, payment := range paymtnts {
		sum += int64(payment.Amount)
	}
	sum /=int64(avg)
	
	return sum
}

func TotalInCategory(paymtnts []types.Payment, category types.Category) int64 {
	sum := int64(0)
	
	for _, pay := range paymtnts {
		
		if pay.Category == category {
			sum += int64(pay.Amount)
		}
	
	}
		
	return sum
}