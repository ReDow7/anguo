package scene

import (
	"anguo/infra/request/tushare"
	"fmt"
	"testing"
)

func TestCompareValueOfAssessmentWithPriceNow(t *testing.T) {
	tushare.InitClient(tushare.GetTokenFromFile())
	date, err := CompareValueOfAssessmentWithPriceNow("603886.SH")
	if err != nil {
		t.Errorf("TestCompareValueOfAssessmentWithPriceNow return error %v", err)
		return
	}
	if date < 2.0 {
		t.Errorf("TestCompareValueOfAssessmentWithPriceNow's result not valid value: %v", date)
		return
	}
	fmt.Println(date)
}
