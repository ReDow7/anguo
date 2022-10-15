package tushare

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLeastCurrentMarketDate(t *testing.T) {
	InitClient("")
	date, err := GetLeastCurrentMarketDate()
	if err != nil {
		t.Errorf("GetLeastCurrentMarketDate return error %v", err)
		return
	}
	res, err := time.Parse("20060102", date)
	if err != nil {
		t.Errorf("GetLeastCurrentMarketDate's result can not be parse %v", err)
		return
	}
	fmt.Println(res)
}
