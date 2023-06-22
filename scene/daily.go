package scene

import (
	"anguo/domain/common"
	"anguo/infra/request/tushare"
	"fmt"
)

func CompareMyHolderOfAssessmentWithPriceDaily() error {
	stocks, err := tushare.ListListingStocks()
	if err != nil {
		return err
	}
	myHoldersCodes := tushare.GetMyHolderCodes()
	historyAssessmentValues := readHistoryFromFile()
	endDate := common.GetLastYearEndDate()
	var myHolders []*CompareResult
	for _, stock := range stocks {
		if !contains(myHoldersCodes, stock.Code) {
			continue
		}
		_, ratio, mv, saturation, err := isCompareRatioMoreThanThreshold(
			historyAssessmentValues, stock.Code, endDate, 0.0)
		if err != nil {
			fmt.Printf("error when comapre on code %s and date %s, %v\n", stock.Code, endDate, err)
			continue
		}
		myHolders = append(myHolders, &CompareResult{
			Stock: stock, Ratio: ratio, PriceValue: mv, saturation: saturation,
		})
	}
	outputDailyCompareResult(myHolders)
	return nil
}

func outputDailyCompareResult(myHolders []*CompareResult) {
	fmt.Println("Code\tRatio\tpriceValue\tName\t")
	for _, result := range myHolders {
		fmt.Printf("%s\t%.2f\t%.2fb\t%s\n", result.Stock.Code, result.Ratio,
			result.PriceValue/1000000000.0, result.Stock.Name)
	}
}

func contains(sli []string, obj string) bool {
	if len(sli) == 0 {
		return false
	}
	for _, e := range sli {
		if e == obj {
			return true
		}
	}
	return false
}
