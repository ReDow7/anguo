package tushare

import (
	"anguo/domain/common"
	"fmt"
	"testing"
)

func TestGetIncomeStatementFromTushare(t *testing.T) {
	InitClient(GetTokenFromFile())
	date, err := GetIncomeStatementFromTushare("000001.SZ")
	if err != nil {
		t.Errorf("GetIncomeStatementFromTushare return error %v", err)
		return
	}
	if date.TotalRevenue < 1.0 {
		t.Errorf("GetIncomeStatementFromTushare's result not valid value: %v", date)
		return
	}
	if date.TotalComprehensiveIncome < 1.0 {
		t.Errorf("GetIncomeStatementFromTushare's result not valid value: %v", date)
		return
	}
	if date.EndDate != common.GetLastYearEndDate() {
		t.Errorf("GetIncomeStatementFromTushare's result not valid value: %v", date)
		return
	}
	fmt.Println(date)
}
