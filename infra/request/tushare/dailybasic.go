package tushare

import (
	"strings"
)

type dailyBasic struct {
	tsCode           string
	tradeDate        string
	totalMarketValue float64
}

func GetTotalMarketValueOfGiveTsCode(tsCode, date string) (float64, error) {
	di, err := getDailyBasicFromTushare(tsCode, date)
	if err != nil {
		return -1e20, err
	}
	return di.totalMarketValue, nil
}

func getDailyBasicFromTushare(tsCode, date string) (*dailyBasic, error) {
	fields := strings.Join([]string{fieldTsCode, fieldTradeDate, fieldTotalMv}, ",")
	params := map[string]interface{}{
		fieldTsCode:    tsCode,
		fieldTradeDate: date,
	}
	resp, err := fetchTushareRawData("daily_basic", fields, params)
	if err != nil {
		return nil, err
	}
	return parseDailyBasicRecord(resp)
}

func parseDailyBasicRecord(resp *Response) (*dailyBasic, error) {
	err := resp.anyError()
	if err != nil {
		return nil, err
	}
	var ret dailyBasic
	for i, field := range resp.Data.Fields {
		switch field {
		case fieldTsCode:
			if str, ok := resp.Data.Items[0][i].(string); ok {
				ret.tsCode = str
			}
		case fieldTotalMv:
			if val, ok := resp.Data.Items[0][i].(float64); ok {
				ret.totalMarketValue = val
			}
		case fieldCalDate:
			if str, ok := resp.Data.Items[0][i].(string); ok {
				ret.tradeDate = str
			}
		default:
			continue
		}
	}
	return &ret, nil
}
