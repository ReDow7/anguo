package tushare

import (
	"fmt"
	"testing"
)

func TestGetCashDividend(t *testing.T) {
	InitClient(GetTokenFromFile())
	date, err := GetCashDividend("603886.SH")
	if err != nil {
		t.Errorf("GetCashDividend return error %v", err)
		return
	}
	if len(date) < 3 || date[0].GetTotalDividendCash() < 1e5 {
		t.Errorf("GetCashDividend's result not valid value: %v", date)
		return
	}
	fmt.Println(date)
}
