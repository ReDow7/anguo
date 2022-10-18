package assessment

import (
	"anguo/infra/request/tushare"
	"fmt"
	"testing"
)

func TestROCEAssessment(t *testing.T) {
	tushare.InitClient(tushare.GetTokenFromFile())
	date, err := ROCEAssessment("603886.SH", "20211231", 0.1)
	if err != nil {
		t.Errorf("ROCEAssessment return error %v", err)
		return
	}
	if date.ValueUnderSustainableGrowthAt4Percent < 2e9 || date.ValueUnderSustainableGrowthAt4Percent > 1e10 {
		t.Errorf("ROCEAssessment's result not valid value: %v", date)
		return
	}
	fmt.Println(date)
}
