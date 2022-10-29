package scene

import (
	"anguo/domain/assessment"
	"anguo/infra/request/tushare"
	"fmt"
	"math"
	"testing"
)

func TestFileSave(t *testing.T) {
	var data = make(map[string]*dataSavedEntry)
	data["1123"] = &dataSavedEntry{
		code:            "1123",
		assessmentValue: 123456.0,
		date:            "20061231",
	}
	data["3211"] = &dataSavedEntry{
		code:            "3211",
		assessmentValue: 654321.0,
		date:            "20071231",
	}
	saveDataToFile(data)
	data = readHistoryFromFile()
	if data["1123"] == nil || data["3211"] == nil {
		t.Errorf("testFileSave can not read data \n")
		return
	}
	if data["1123"].code != "1123" || math.Abs(data["1123"].assessmentValue-123456.0) > 1e-2 ||
		data["1123"].date != "20061231" {
		t.Errorf("testFileSave data read error \n")
		return
	}
}

func TestCompareValueOfAssessmentWithAllStock(t *testing.T) {
	tushare.InitClient(tushare.GetTokenFromFile())
	result, err := CompareAllStockValueOfAssessmentWithPriceNow(0, 3)
	if err != nil {
		t.Errorf("CompareAllStockValueOfAssessmentWithPriceNow return error\n")
		return
	}
	if len(result) != 3 {
		t.Errorf("CompareAllStockValueOfAssessmentWithPriceNow return data is not expected\n")
		return
	}
	tempFunc := assessmentFunc
	assessmentFunc = func(code string, date string, WACC float64) (*assessment.ROCEAssessmentResult, error) {
		return nil, fmt.Errorf("intend error for unit test")
	}
	result, err = CompareAllStockValueOfAssessmentWithPriceNow(0, 3)
	if err != nil {
		t.Errorf("CompareAllStockValueOfAssessmentWithPriceNow return error after first time\n")
		return
	}
	if len(result) != 3 {
		t.Errorf("CompareAllStockValueOfAssessmentWithPriceNow return data is not expected after first time\n")
		return
	}
	assessmentFunc = tempFunc
}

func TestCompareValueOfAssessmentWithPriceNow(t *testing.T) {
	tushare.InitClient(tushare.GetTokenFromFile())
	_, _, date, err := CompareValueOfAssessmentWithPriceNow("603886.SH", nil)
	if err != nil {
		t.Errorf("TestCompareValueOfAssessmentWithPriceNow return error %v", err)
		return
	}
	if *date < 2.0 {
		t.Errorf("TestCompareValueOfAssessmentWithPriceNow's result not valid value: %v", *date)
		return
	}
	fmt.Println(*date)
}
