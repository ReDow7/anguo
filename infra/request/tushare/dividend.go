package tushare

import (
	"anguo/model"
	"strings"
)

func GetCashDividend(tsCode string) ([]*model.CashDividend, error) {
	di, err := getCashDividendFromTushare(tsCode)
	if err != nil {
		return nil, err
	}
	return di, nil
}

func getCashDividendFromTushare(tsCode string) ([]*model.CashDividend, error) {
	fields := strings.Join([]string{fieldTsCode, fieldEndDate, fieldDivProc, fieldCashDivTax, fieldBaseShare}, ",")
	params := map[string]interface{}{
		fieldTsCode: tsCode,
	}
	resp, err := fetchTushareRawData("dividend", fields, params)
	if err != nil {
		return nil, err
	}
	return parseCashDividendRecord(resp)
}

func parseCashDividendRecord(resp *Response) ([]*model.CashDividend, error) {
	err := resp.anyError()
	if err != nil {
		return nil, err
	}
	var dividends []*model.CashDividend
	var dateMark = make(map[string]bool)
	for _, item := range resp.Data.Items {
		var div model.CashDividend
		var effect = false
		for i, field := range resp.Data.Fields {
			switch field {
			case fieldTsCode:
				if str, ok := item[i].(string); ok {
					div.Code = str
				}
			case fieldEndDate:
				if str, ok := item[i].(string); ok {
					div.Date = str
				}
			case fieldDivProc:
				if str, ok := item[i].(string); ok {
					if str == "实施" {
						effect = true
					}
				}
			case fieldBaseShare:
				if val, ok := item[i].(float64); ok {
					div.BaseShare = val
				}
			case fieldCashDivTax:
				if val, ok := item[i].(float64); ok {
					div.CashDividendPerShare = val
				}
			default:
				continue
			}
		}
		if effect && !dateMark[div.Date] {
			dateMark[div.Date] = true
			dividends = append(dividends, &div)
		}
	}
	return dividends, nil
}
