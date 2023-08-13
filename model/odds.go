package model

import (
	"fmt"
	"math"
)

func CalculateOdds(assertOfFinancial, evaluate, price float64) (odds float64, desc string) {
	avg, sgm := evaluate, (evaluate-assertOfFinancial)/3.0
	if sgm <= 1e-8 {
		return evaluate / price, fmt.Sprintf("%.2f", evaluate/price)
	}
	start, end, result := avg-3*sgm, avg+3*sgm, 0.0
	step := (end - start) / 1000.0
	last := 0.0
	for i := start; i <= end; i += step {
		pro := normalDistribution(i/1e9, avg/1e9, sgm/1e9)
		result += (i - price) * math.Abs(pro-last)
		last = pro
	}
	return result / price, fmt.Sprintf("%.2f", result/price)
}

func normalDistribution(per float64, avg float64, sgm float64) float64 {
	return math.Pow(math.E, -((per-avg)*(per-avg))/(2*sgm*sgm)) / (sgm * math.Sqrt(2.0*math.Pi))
}
