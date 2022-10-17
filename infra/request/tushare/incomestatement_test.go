package tushare

import (
	"fmt"
	"testing"
)

func TestGetIncomeStatementFromTushare(t *testing.T) {
	InitClient(GetTokenFromFile())
	date, err := GetIncomeStatementFromTushare("603886.SH")
	if err != nil {
		t.Errorf("GetIncomeStatementFromTushare return error %v", err)
		return
	}
	if date.TotalRevenue < 1.0 {
		t.Errorf("GetIncomeStatementFromTushare's result not valid value: %v", date)
		return
	}
	fmt.Println(date)
}
