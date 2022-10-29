package scene

import (
	"anguo/infra/request/tushare"
	"fmt"
	"time"
)

func GetThreeYearsAverageDividendForCodeAndDateGive(code, date string) (*float64, error) {
	thisYear, err := time.Parse("20060102", date)
	if err != nil {
		return nil, fmt.Errorf("error when parse time of thisYear val: %s\n", thisYear)
	}
	threeYearsAgo := thisYear.AddDate(-3, 0, 0)
	dividend, err := tushare.GetCashDividend(code)
	if err != nil {
		return nil, fmt.Errorf("can not get cash dividend code : %s thisYear: %s\n", code, thisYear)
	}
	totalCashAmount := 0.0
	for _, d := range dividend {
		dividendTime, err := time.Parse("20060102", d.Date)
		if err != nil {
			return nil, err
		}
		if dividendTime.After(threeYearsAgo) {
			totalCashAmount += d.GetTotalDividendCash()
		}
	}
	ret := totalCashAmount / 3.0
	return &ret, nil
}
