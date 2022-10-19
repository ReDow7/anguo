package scene

import (
	"anguo/domain/assessment"
	"anguo/domain/common"
	"anguo/infra/request/tushare"
	"fmt"
)

const (
	AVERAGE_WACC = 0.1
)

func CompareValueOfAssessmentWithPriceNow(tsCode string) (float64, error) {
	assessmentValue, error := assessment.ROCEAssessment(tsCode, common.GetLastYearEndDate(), AVERAGE_WACC)
	if error != nil {
		return -1e20, error
	}
	lastMarketDay, error := tushare.GetLeastCurrentMarketDate()
	if error != nil {
		return -1e20, error
	}
	price, error := tushare.GetTotalMarketValueOfGiveTsCode(tsCode, lastMarketDay)
	if error != nil {
		return -1e20, error
	}
	fmt.Printf("CompareValueOfAssessmentWithPriceNow -> tsCode : %s price %f vs assement value %f ",
		tsCode, price, assessmentValue.ValueUnderSustainableGrowthAt6Percent)
	return assessmentValue.ValueUnderSustainableGrowthAt6Percent / price, nil
}
