package scene

import (
	"anguo/domain/assessment"
	"anguo/domain/common"
	"anguo/infra/dal"
	"anguo/infra/request/tushare"
	"anguo/model"
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	averageWACC           = 0.1
	dataFileName          = "compare.sec"
	compareResultFileName = "compare.result.sec"
)

type dataSavedEntry struct {
	code            string
	assessmentValue float64
	date            string
}

type CompareResult struct {
	Stock      model.Stock
	Ratio      float64
	PriceValue float64
}

type historyCompareResult struct {
	code         string
	compareRatio float64
}

var assessmentFunc = assessment.ROCEAssessment

func readHistoryFromFile() map[string]*dataSavedEntry {
	var ret = make(map[string]*dataSavedEntry)
	saved, err := dal.ReadFromFile(dataFileName)
	if err != nil {
		_ = fmt.Errorf("can not read from file with error : %v\n", err)
		return ret
	}
	lines := strings.Split(saved, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			_ = fmt.Errorf("a valid line from file : %s\n", line)
			continue
		}
		value, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			_ = fmt.Errorf("a valid value line from file : %s\n", line)
			continue
		}
		ret[parts[0]] = &dataSavedEntry{
			code: parts[0], date: parts[1], assessmentValue: value,
		}
	}
	fmt.Printf("%d lines read from file successfully\n", len(ret))
	return ret
}

func saveDataToFile(data map[string]*dataSavedEntry) {
	if len(data) == 0 {
		_ = fmt.Errorf("no date to write, return directly\n")
		return
	}
	var buf bytes.Buffer
	for _, entry := range data {
		buf.WriteString(strings.Join([]string{entry.code, entry.date,
			fmt.Sprintf("%.2f", entry.assessmentValue)}, ","))
		buf.WriteString("\n")
	}
	err := dal.WriteToFileOverWrite(dataFileName, buf.String())
	if err != nil {
		_ = fmt.Errorf("can not write to file with error : %v\n", err)
		return
	}
	fmt.Printf("write to file success data entries : %d\n", len(data))
}

func CompareAllStockValueOfAssessmentWithPriceNow(compareThreshold float64, numberLimit int) ([]*CompareResult, error) {
	stocks, err := tushare.ListListingStocks()
	if err != nil {
		return nil, err
	}
	historyAssessmentValues := readHistoryFromFile()
	endDate := common.GetLastYearEndDate()
	var picks []*CompareResult
	for i, stock := range stocks {
		if i == numberLimit {
			fmt.Printf("touch number limit %d\n", numberLimit)
			break
		}
		if !hasListLongThanThreeYears(&stock) {
			continue
		}
		pick, ratio, err := isCompareRatioMoreThanThreshold(
			historyAssessmentValues, stock.Code, endDate, compareThreshold)
		if err != nil {
			_ = fmt.Errorf("error when comapre on code %s and date %s, %v\n", stock.Code, endDate, err)
			continue
		}
		if pick {
			picks = append(picks, &CompareResult{
				stock, ratio, historyAssessmentValues[stock.Code].assessmentValue,
			})
		}
	}
	saveDataToFile(historyAssessmentValues)
	outputCompareResult(picks)
	return picks, nil
}

func outputCompareResult(results []*CompareResult) {
	if len(results) <= 0 {
		fmt.Println("--NO COMPARE RESULT THIS TIME--")
		return
	}
	history := readHistoryCompareResultFromFile()
	listBefore := make(map[string]bool)
	for _, saved := range history {
		listBefore[saved.code] = true
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Ratio > results[j].Ratio
	})
	fmt.Println("Code\tName\tRatio\tpriceValue\tIndustry")
	for _, result := range results {
		if !listBefore[result.Stock.Code] {
			continue
		}
		fmt.Printf("%s\t%s\t%.2f\t%.2fm\t%s\n", result.Stock.Code, result.Stock.Name, result.Ratio,
			result.PriceValue/1000000.0, result.Stock.Industry)
	}
	fmt.Println("--NEW LIST OF THIS TIME--")
	for _, result := range results {
		if listBefore[result.Stock.Code] {
			continue
		}
		fmt.Printf("%s\t%s\t%.2f\t%.2fm\t%s\n", result.Stock.Code, result.Stock.Name, result.Ratio,
			result.PriceValue/1000000.0, result.Stock.Industry)
	}
	saveCompareResultToFile(results)
}

func saveCompareResultToFile(data []*CompareResult) {
	if len(data) == 0 {
		_ = fmt.Errorf("no compare result need to write, return directly\n")
		return
	}
	var buf bytes.Buffer
	for _, entry := range data {
		buf.WriteString(strings.Join([]string{entry.Stock.Code, fmt.Sprintf("%.2f", entry.Ratio)}, ","))
		buf.WriteString("\n")
	}
	err := dal.WriteToFileOverWrite(compareResultFileName, buf.String())
	if err != nil {
		_ = fmt.Errorf("can not write to file for compare result with error : %v\n", err)
		return
	}
	fmt.Printf("write to file success compare results : %d\n", len(data))
}

func readHistoryCompareResultFromFile() []*historyCompareResult {
	var ret = make([]*historyCompareResult, 0)
	saved, err := dal.ReadFromFile(compareResultFileName)
	if err != nil {
		_ = fmt.Errorf("can not read compare result from file with error : %v\n", err)
		return ret
	}
	lines := strings.Split(saved, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			_ = fmt.Errorf("a valid line from file : %s\n", line)
			continue
		}
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			_ = fmt.Errorf("a valid value line from file : %s\n", line)
			continue
		}
		ret = append(ret, &historyCompareResult{
			parts[0], value,
		})
	}
	fmt.Printf("%d lines read from compare history file successfully\n", len(ret))
	return ret
}

func isCompareRatioMoreThanThreshold(historyAssessmentValues map[string]*dataSavedEntry,
	tsCode, endDate string, threshold float64) (
	bool, float64, error) {
	var err error
	var ratio *float64
	needCalAssessment := true
	if saveEntry, ok := historyAssessmentValues[tsCode]; ok {
		if saveEntry.date == endDate {
			_, _, ratio, err = CompareValueOfAssessmentWithPriceNow(tsCode, &saveEntry.assessmentValue)
			needCalAssessment = false
		}
	}
	if needCalAssessment {
		var value *float64
		value, _, ratio, err = CompareValueOfAssessmentWithPriceNow(tsCode, nil)
		historyAssessmentValues[tsCode] = &dataSavedEntry{
			code:            tsCode,
			date:            endDate,
			assessmentValue: *value,
		}
	}
	if err != nil {
		return false, 0, err
	}
	return *ratio > threshold, *ratio, nil
}

func hasListLongThanThreeYears(stock *model.Stock) bool {
	listTime, err := time.Parse(stock.ListingDate, "20060102")
	if err != nil {
		_ = fmt.Errorf("encouter a not valid list date %s\n", stock.ListingDate)
		return true
	}
	now := time.Now()
	threeYearsAgo := now.AddDate(-3, 0, 0)
	return listTime.Before(threeYearsAgo)
}

func CompareValueOfAssessmentWithPriceNow(tsCode string, assessmentValueGiven *float64) (
	assessmentValue, marketPrice, compareRatio *float64, err error) {
	var assessmentValues *assessment.ROCEAssessmentResult
	if assessmentValueGiven != nil {
		assessmentValues = &assessment.ROCEAssessmentResult{
			ValueUnderSustainableGrowthAt4Percent: *assessmentValueGiven,
			ValueUnderSustainableGrowthAt6Percent: *assessmentValueGiven,
			ValueUnderSustainableGrowthAt8Percent: *assessmentValueGiven,
		}
	} else {
		assessmentValues, err = assessmentFunc(tsCode, common.GetLastYearEndDate(), averageWACC)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	lastMarketDay, err := tushare.GetLeastCurrentMarketDate()
	if err != nil {
		return nil, nil, nil, err
	}
	price, err := tushare.GetTotalMarketValueOfGiveTsCode(tsCode, lastMarketDay)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Printf("CompareValueOfAssessmentWithPriceNow -> tsCode : %s price %f vs assement value %f \n",
		tsCode, price, assessmentValues.ValueUnderSustainableGrowthAt6Percent)
	ratio := assessmentValues.ValueUnderSustainableGrowthAt6Percent / price
	return &assessmentValues.ValueUnderSustainableGrowthAt6Percent, &price, &ratio, nil
}
