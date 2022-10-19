package model

type CashDividend struct {
	Code                 string
	Date                 string
	BaseShare            float64
	CashDividendPerShare float64
}

func (c *CashDividend) GetTotalDividendCash() float64 {
	return c.BaseShare * 10000 * c.CashDividendPerShare
}
