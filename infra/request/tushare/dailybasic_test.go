package tushare

import (
	"fmt"
	"testing"
)

func TestGetTotalMarketValueOfGiveTsCode(t *testing.T) {
	InitClient(GetTokenFromFile())
	date, err := GetTotalMarketValueOfGiveTsCode("603886.SH", "20221014")
	if err != nil {
		t.Errorf("GetTotalMarketValueOfGiveTsCode return error %v", err)
		return
	}
	if *date < 1.0 {
		t.Errorf("GetTotalMarketValueOfGiveTsCode's result not valid value: %v", date)
		return
	}
	fmt.Println(*date)
}
