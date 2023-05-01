package assessment

import (
	"anguo/infra/request/tushare"
)

type ROCEAssessmentResult struct {
	ValueUnderSustainableGrowthAt4Percent float64
	ValueUnderSustainableGrowthAt6Percent float64
	ValueUnderSustainableGrowthAt8Percent float64
	Saturation                            float64
}

func ROCEAssessment(code string, date string, WACC float64) (*ROCEAssessmentResult, error) {
	statement, err := tushare.GetIncomeStatementFromTushareForGivenCodeAndDate(code, date)
	if err != nil {
		return nil, err
	}
	sheet, err := tushare.GetBalanceSheetForGiveCodeAndDate(code, date)
	if err != nil {
		return nil, err
	}
	var value ROCEAssessmentResult
	re := statement.RE(sheet, WACC)
	if re < 0 {
		re = 0
	}
	value.ValueUnderSustainableGrowthAt4Percent = sheet.TotalEquityOfOwnersExcludeMinorityInterests +
		re/(WACC-0.04) - sheet.MinorityInterests
	value.ValueUnderSustainableGrowthAt6Percent = sheet.TotalEquityOfOwnersExcludeMinorityInterests +
		re/(WACC-0.06) - sheet.MinorityInterests
	value.ValueUnderSustainableGrowthAt8Percent = sheet.TotalEquityOfOwnersExcludeMinorityInterests +
		re/(WACC-0.08) - sheet.MinorityInterests

	value.Saturation = sheet.NetFinancialAssert(statement.Revenue) / value.ValueUnderSustainableGrowthAt6Percent
	return &value, nil
}
