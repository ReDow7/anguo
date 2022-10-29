package tushare

import (
	"fmt"
	"testing"
)

func TestGetStockList(t *testing.T) {
	InitClient(GetTokenFromFile())
	stocks, err := getStockListFromTushare()
	if err != nil {
		t.Errorf("getStockListFromTushare return error %v", err)
		return
	}
	if len(stocks) < 1000 {
		t.Errorf("getStockListFromTushare return too little stocks size: %d", len(stocks))
		return
	}
	fmt.Println(stocks[3])
}
