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

var OutputFile = ""

type dataSavedEntry struct {
	code            string
	assessmentValue float64
	date            string
	saturation      float64
}

type CompareResult struct {
	Stock               model.Stock
	Ratio               float64
	PriceValue          float64
	averageDividendRate float64
	alerts              []AlertInfo
	saturation          float64
}

type historyCompareResult struct {
	code                string
	name                string
	appear              bool
	compareRatio        float64
	averageDividendRate float64
	alertInfo           []AlertInfo
}

var assessmentFunc = assessment.ROCEAssessment

func readHistoryFromFile() map[string]*dataSavedEntry {
	var ret = make(map[string]*dataSavedEntry)
	saved, err := dal.ReadFromFile(dataFileName)
	if err != nil {
		fmt.Printf("can not read from file with error : %v\n", err)
		return ret
	}
	lines := strings.Split(saved, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			fmt.Printf("a valid line from file : %s\n", line)
			continue
		}
		value, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Printf("a valid value line from file : %s\n", line)
			continue
		}
		saturation, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			fmt.Printf("a valid saturation line from file : %s\n", line)
			continue
		}
		ret[parts[0]] = &dataSavedEntry{
			code: parts[0], date: parts[1], assessmentValue: value, saturation: saturation,
		}
	}
	fmt.Printf("%d lines read from file successfully\n", len(ret))
	return ret
}

func saveDataToFile(data map[string]*dataSavedEntry) {
	if len(data) == 0 {
		fmt.Printf("no date to write, return directly\n")
		return
	}
	var buf bytes.Buffer
	for _, entry := range data {
		buf.WriteString(strings.Join([]string{entry.code, entry.date,
			fmt.Sprintf("%.2f,%.2f", entry.assessmentValue, entry.saturation)}, ","))
		buf.WriteString("\n")
	}
	err := dal.WriteToFileOverWrite(dataFileName, buf.String())
	if err != nil {
		fmt.Printf("can not write to file with error : %v\n", err)
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
	var myHolders []*CompareResult
	var codeToName = make(map[string]string)
	myHoldersCodes := tushare.GetMyHolderCodes()
	for i, stock := range stocks {
		fmt.Printf("compare process %d/%d\n", i, len(stocks))
		codeToName[stock.Code] = stock.Name
		if i == numberLimit {
			fmt.Printf("touch number limit %d\n", numberLimit)
			break
		}
		if !hasListLongThanThreeYears(&stock) {
			continue
		}
		pick, ratio, mv, saturation, err := isCompareRatioMoreThanThreshold(
			historyAssessmentValues, stock.Code, endDate, compareThreshold)
		if err != nil {
			fmt.Printf("error when comapre on code %s and date %s, %v\n", stock.Code, endDate, err)
			continue
		}
		if pick {
			picks = append(picks, &CompareResult{
				Stock: stock, Ratio: ratio, PriceValue: mv, saturation: saturation,
			})
		}
		if isInMyHolders(stock.Code, myHoldersCodes) {
			myHolders = append(myHolders, &CompareResult{
				Stock: stock, Ratio: ratio, PriceValue: mv, saturation: saturation,
			})
		}
	}
	listBefore := make(map[string]*historyCompareResult)
	history := readHistoryCompareResultFromFile()
	for _, saved := range history {
		listBefore[saved.code] = saved
	}
	for _, e := range listBefore {
		e.name = codeToName[e.code]
	}
	saveDataToFile(historyAssessmentValues)
	outputCompareResult(myHolders, picks, endDate, listBefore)
	return picks, nil
}

func outputCompareResult(myHolders, results []*CompareResult, endDate string, listBefore map[string]*historyCompareResult) {
	var buf bytes.Buffer
	sort.Slice(results, func(i, j int) bool {
		return results[i].Ratio > results[j].Ratio
	})
	if outputToFile() {
		buf.WriteString("Code\tName\tRatio\tpriceValue\tIndustry\tDividend\tAlertInfo\tsaturation\n")
	}

	fmt.Println("Code\tName\tRatio\tpriceValue\tIndustry\tDividend\tAlertInfo\tsaturation")
	for _, result := range results {
		if listBefore[result.Stock.Code] == nil {
			continue
		}
		saved := listBefore[result.Stock.Code]
		if saved.averageDividendRate < -1e-8 {
			saved.averageDividendRate = collectAverageDividendRate(result.Stock.Code, endDate)
		}
		if len(saved.alertInfo) == 0 || saved.alertInfo[0] == alertError {
			saved.alertInfo = CollectAlertInfosForCodeAndDataGive(result.Stock.Code, endDate)
		}
		result.averageDividendRate = saved.averageDividendRate
		result.alerts = saved.alertInfo
		if outputToFile() {
			buf.WriteString(fmt.Sprintf("%s\t%s\t%.2f\t%.2fm\t%s\t%.2f\t%s\t%.2f\n", result.Stock.Code, result.Stock.Name, result.Ratio,
				result.PriceValue/1000000.0, result.Stock.Industry,
				result.averageDividendRate, generateAlertInfoToSave(result.alerts), result.saturation))
		}
		fmt.Printf("%s\t%s\t%.2f\t%.2fm\t%s\t%.2f\t%s\t%.2f\n", result.Stock.Code, result.Stock.Name, result.Ratio,
			result.PriceValue/1000000.0, result.Stock.Industry,
			result.averageDividendRate, generateAlertInfoToSave(result.alerts), result.saturation)
	}
	if outputToFile() {
		buf.WriteString("--NEW LIST OF THIS TIME--\n")
	}
	fmt.Println("--NEW LIST OF THIS TIME--")
	for _, result := range results {
		if listBefore[result.Stock.Code] != nil {
			listBefore[result.Stock.Code].appear = true
			continue
		}
		dividendRate := collectAverageDividendRate(result.Stock.Code, endDate)
		alerts := CollectAlertInfosForCodeAndDataGive(result.Stock.Code, endDate)
		result.averageDividendRate = dividendRate
		result.alerts = alerts
		if outputToFile() {
			buf.WriteString(fmt.Sprintf("%s\t%s\t%.2f\t%.2fm\t%s\t%.2f\t%s\t%.2f\n", result.Stock.Code, result.Stock.Name, result.Ratio,
				result.PriceValue/1000000.0, result.Stock.Industry,
				result.averageDividendRate, generateAlertInfoToSave(result.alerts), result.saturation))
		}
		fmt.Printf("%s\t%s\t%.2f\t%.2fm\t%s\t%.2f\t%s\t%.2f\n", result.Stock.Code, result.Stock.Name, result.Ratio,
			result.PriceValue/1000000.0, result.Stock.Industry,
			result.averageDividendRate, generateAlertInfoToSave(result.alerts), result.saturation)
	}
	fmt.Println("--NOT SHOWED THIS TIME--")
	for _, result := range listBefore {
		if result.appear {
			continue
		}
		if outputToFile() {
			buf.WriteString(fmt.Sprintf("%s\t%s\n", result.code, result.name))
		}
		fmt.Printf("%s\t%s\n", result.code, result.name)
	}
	fmt.Println("--MY HOLDER--")
	for _, result := range myHolders {
		if listBefore[result.Stock.Code] != nil {
			continue
		}
		dividendRate := collectAverageDividendRate(result.Stock.Code, endDate)
		alerts := CollectAlertInfosForCodeAndDataGive(result.Stock.Code, endDate)
		result.averageDividendRate = dividendRate
		result.alerts = alerts
		if outputToFile() {
			buf.WriteString(fmt.Sprintf("%s\t%s\t%.2f\t%.2fm\t%s\t%.2f\t%s\t%.2f\n", result.Stock.Code, result.Stock.Name, result.Ratio,
				result.PriceValue/1000000.0, result.Stock.Industry,
				result.averageDividendRate, generateAlertInfoToSave(result.alerts), result.saturation))
		}
		fmt.Printf("%s\t%s\t%.2f\t%.2fm\t%s\t%.2f\t%s\t%.2f\n", result.Stock.Code, result.Stock.Name, result.Ratio,
			result.PriceValue/1000000.0, result.Stock.Industry,
			result.averageDividendRate, generateAlertInfoToSave(result.alerts), result.saturation)
	}
	saveCompareResultToFile(results)
	if outputToFile() {
		_ = dal.WriteToFileOverWrite(OutputFile, buf.String())
	}
}

func outputToFile() bool {
	return len(OutputFile) > 0
}

func collectAverageDividendRate(code, endDate string) float64 {
	averageDividendLast3Years, err := GetThreeYearsAverageDividendForCodeAndDateGive(code, endDate)
	if err != nil {
		fmt.Printf("error when get last 3 years cash dividen for code %s date %s",
			code, endDate)
	}
	var dividendRate = -1.0
	if averageDividendLast3Years == nil {
		return dividendRate
	}
	lastMarketDay, err := tushare.GetLeastCurrentMarketDate()
	if err != nil {
		fmt.Printf("error when get last market day %s\n", err)
		return dividendRate
	}
	mv, err := tushare.GetTotalMarketValueOfGiveTsCode(code, lastMarketDay)
	if err != nil {
		fmt.Printf("error when get total market value %s\n", err)
		return dividendRate
	}
	dividendRate = *averageDividendLast3Years / *mv
	return dividendRate
}

func saveCompareResultToFile(data []*CompareResult) {
	if len(data) == 0 {
		fmt.Printf("no compare result need to write, return directly\n")
		return
	}
	var buf bytes.Buffer
	for _, entry := range data {
		buf.WriteString(strings.Join([]string{
			entry.Stock.Code, fmt.Sprintf("%.2f", entry.Ratio),
			fmt.Sprintf("%.2f", entry.averageDividendRate), generateAlertInfoToSave(entry.alerts)}, ","))
		buf.WriteString("\n")
	}
	err := dal.WriteToFileOverWrite(compareResultFileName, buf.String())
	if err != nil {
		fmt.Printf("can not write to file for compare result with error : %v\n", err)
		return
	}
	fmt.Printf("write to file success compare results : %d\n", len(data))
}

func readHistoryCompareResultFromFile() []*historyCompareResult {
	var ret = make([]*historyCompareResult, 0)
	saved, err := dal.ReadFromFile(compareResultFileName)
	if err != nil {
		fmt.Printf("can not read compare result from file with error : %v\n", err)
		return ret
	}
	lines := strings.Split(saved, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			fmt.Printf("a valid line from file : %s\n", line)
			continue
		}
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Printf("a valid value line from file : %s\n", line)
			continue
		}
		averageDividendRate, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Printf("a valid value line from file : %s\n", line)
			continue
		}
		ret = append(ret, &historyCompareResult{
			code:                parts[0],
			compareRatio:        value,
			averageDividendRate: averageDividendRate,
			alertInfo:           parseAlertInfo(parts[3]),
		})
	}
	fmt.Printf("%d lines read from compare history file successfully\n", len(ret))
	return ret
}

func parseAlertInfo(str string) []AlertInfo {
	ret := make([]AlertInfo, 0)
	if len(str) == 0 {
		return ret
	}
	segs := strings.Split(str, "_")
	for _, info := range segs {
		val, err := strconv.Atoi(info)
		if err != nil {
			fmt.Printf("a valid val read from file, %s\n", str)
			continue
		}
		ret = append(ret, AlertInfo(val))
	}
	return ret
}

func generateAlertInfoToSave(infos []AlertInfo) string {
	if len(infos) == 0 {
		return fmt.Sprintf("%d", noAlert)
	}
	strList := make([]string, 0)
	for _, info := range infos {
		strList = append(strList, fmt.Sprintf("%d", info))
	}
	return strings.Join(strList, "_")
}

func isCompareRatioMoreThanThreshold(historyAssessmentValues map[string]*dataSavedEntry,
	tsCode, endDate string, threshold float64) (
	bool, float64, float64, float64, error) {
	var err error
	var ratio, mv, saturation *float64
	needCalAssessment := true
	if saveEntry, ok := historyAssessmentValues[tsCode]; ok {
		if saveEntry.date == endDate {
			_, mv, ratio, saturation, err = CompareValueOfAssessmentWithPriceNow(tsCode, &saveEntry.assessmentValue)
			needCalAssessment = false
		}
	}
	if needCalAssessment {
		var value *float64
		value, mv, ratio, saturation, err = CompareValueOfAssessmentWithPriceNow(tsCode, nil)
		if err != nil {
			return false, 0, 0, 0, err
		}
		historyAssessmentValues[tsCode] = &dataSavedEntry{
			code:            tsCode,
			date:            endDate,
			assessmentValue: *value,
		}
	}
	if err != nil {
		return false, 0, 0, 0, err
	}
	return *ratio > threshold, *ratio, *mv, *saturation, nil
}

func hasListLongThanThreeYears(stock *model.Stock) bool {
	listTime, err := time.Parse("20060102", stock.ListingDate)
	if err != nil {
		fmt.Printf("encouter a not valid list date %s\n", stock.ListingDate)
		return true
	}
	now := time.Now()
	threeYearsAgo := now.AddDate(-3, 0, 0)
	return listTime.Before(threeYearsAgo)
}

func CompareValueOfAssessmentWithPriceNow(tsCode string, assessmentValueGiven *float64) (
	assessmentValue, marketPrice, compareRatio *float64, saturation *float64, err error) {
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
			return nil, nil, nil, nil, err
		}
	}
	lastMarketDay, err := tushare.GetLeastCurrentMarketDate()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	price, err := tushare.GetTotalMarketValueOfGiveTsCode(tsCode, lastMarketDay)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	fmt.Printf("CompareValueOfAssessmentWithPriceNow -> tsCode : %s price %f vs assement value %f \n",
		tsCode, (*price)/100000000, assessmentValues.ValueUnderSustainableGrowthAt6Percent/100000000)
	ratio := assessmentValues.ValueUnderSustainableGrowthAt6Percent / *price
	return &assessmentValues.ValueUnderSustainableGrowthAt6Percent, price, &ratio, &assessmentValues.Saturation, nil
}

func isInMyHolders(code string, codes []string) bool {
	for _, str := range codes {
		if code == str {
			return true
		}
	}
	return false
}
