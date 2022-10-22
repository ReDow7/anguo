package scene

import (
	"anguo/infra/request/tushare"
	"fmt"
	"time"
)

type AlertInfo int

const (
	alertError          = AlertInfo(-1)
	noAlert             = AlertInfo(0)
	highIncomeLastYear  = AlertInfo(1)
	hasLossInThreeYears = AlertInfo(2)
)

func CollectAlertInfosForCodeAndDataGive(code, date string) []AlertInfo {
	ret := make([]AlertInfo, 0)
	ciOfLast3years, err := getLast3YearsTotalComprehensiveIncome(code, date)
	if err != nil {
		return []AlertInfo{alertError}
	}
	if hasHighIncomeLastYearAlert(ciOfLast3years) {
		ret = append(ret, highIncomeLastYear)
	}
	if hasHasLossInThreeYearsAlert(ciOfLast3years) {
		ret = append(ret, hasLossInThreeYears)
	}
	if len(ret) == 0 {
		ret = append(ret, noAlert)
	}
	return ret
}

func hasHighIncomeLastYearAlert(last3Years []float64) bool {
	average := (last3Years[0] + last3Years[1] + last3Years[2]) / 3
	return last3Years[0] > average*2
}

func hasHasLossInThreeYearsAlert(last3Years []float64) bool {
	for _, val := range last3Years {
		if val < 0 {
			return true
		}
	}
	return false
}

func getLast3YearsTotalComprehensiveIncome(code, date string) ([]float64, error) {
	ret := make([]float64, 3)
	thisYear, err := time.Parse("20060102", date)
	if err != nil {
		return nil, fmt.Errorf("error when parse time of thisYear val: %s\n", thisYear)
	}
	statement, err := tushare.GetIncomeStatementFromTushareForGivenCodeAndDate(
		code, thisYear.Format("20060102"))
	if err != nil {
		return nil, err
	}
	ret = append(ret, statement.TotalComprehensiveIncome)
	lastYear := thisYear.AddDate(-1, 0, 0)
	statement, err = tushare.GetIncomeStatementFromTushareForGivenCodeAndDate(
		code, lastYear.Format("20060102"))
	if err != nil {
		return nil, err
	}
	ret = append(ret, statement.TotalComprehensiveIncome)
	lastTwoYear := thisYear.AddDate(-2, 0, 0)
	statement, err = tushare.GetIncomeStatementFromTushareForGivenCodeAndDate(
		code, lastTwoYear.Format("20060102"))
	if err != nil {
		return nil, err
	}
	ret = append(ret, statement.TotalComprehensiveIncome)
	return ret, nil
}
