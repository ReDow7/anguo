package tushare

import (
	"fmt"
	"testing"
)

func TestGetBalanceSheetOfLastYearForGiveCode(t *testing.T) {
	InitClient(GetTokenFromFile())
	date, err := GetBalanceSheetOfLastYearForGiveCode("603886.SH")
	if err != nil {
		t.Errorf("GetBalanceSheetOfLastYearForGiveCode return error %v", err)
		return
	}
	if date.TotalCurrentAsserts < 1.0 {
		t.Errorf("GetBalanceSheetOfLastYearForGiveCode's result not valid value: %v", date)
		return
	}
	fmt.Println(date)
}
