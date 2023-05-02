package tushare

import (
	"testing"
)

func TestMyHolderList(t *testing.T) {
	list := GetMyHolderCodes()
	if !isIn("1231.SZ", list) || !isIn("1232.ZH", list) {
		t.Errorf("can not read correct my holder list")
	}
}

func isIn(code string, codes []string) bool {
	for _, str := range codes {
		if code == str {
			return true
		}
	}
	return false
}
