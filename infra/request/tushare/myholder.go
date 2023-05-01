package tushare

import (
	"anguo/infra/dal"
	"fmt"
	"strings"
)

func GetMyHolderCodes() []string {
	content, err := dal.ReadFromFile("myholder.sec")
	if err != nil {
		fmt.Printf("can not read from myholderfile return emtpy list")
		return make([]string, 0)
	}
	return strings.Split(content, "\n")
}
