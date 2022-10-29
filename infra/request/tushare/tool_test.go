package tushare

import (
	"fmt"
	"testing"
	"time"
)

func TestReadTokenFromFile(t *testing.T) {
	token := GetTokenFromFile()
	if token == "" {
		t.Errorf("can not read token from file")
	}
}

func TestGetLeastCurrentMarketDate(t *testing.T) {
	InitClient(GetTokenFromFile())
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
