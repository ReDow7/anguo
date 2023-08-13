package scene

import (
	"anguo/domain/common"
	"anguo/model"
	"fmt"
)

func CompareOneOfAssessmentWithPriceDaily(code string) error {
	if len(code) == 0 {
		return fmt.Errorf("empty code")
	}
	historyAssessmentValues := readHistoryFromFile()
	endDate := common.GetLastYearEndDate()
	var result *CompareResult

	_, ratio, mv, saturation, err := isCompareRatioMoreThanThreshold(
		historyAssessmentValues, code, endDate, 0.0)
	evaluate := mv * ratio
	odds, _ := model.CalculateOdds(evaluate*saturation, evaluate, mv)
	if err != nil {
		fmt.Printf("error when comapre on code %s and date %s, %v\n", code, endDate, err)
		return err
	}
	result = &CompareResult{
		Ratio: ratio, PriceValue: mv, saturation: saturation, odds: odds,
	}

	outputOneCompareResult(code, result)
	return nil
}

func outputOneCompareResult(code string, result *CompareResult) {
	fmt.Println("Code\tRatio\tpriceValue\todds\t")
	fmt.Printf("%s\t%.2f\t%.2fb\t%.2f\n", code, result.Ratio,
		result.PriceValue/1000000000.0, result.odds)
}
