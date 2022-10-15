package tushare

import (
	"anguo/model"
	"strings"
)

func ListListingStocks() ([]model.Stock, error) {
	stocks, err := getStockListFromTushare()
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func getStockListFromTushare() ([]model.Stock, error) {
	fields := strings.Join([]string{fieldTsCode, fieldName, fieldArea, fieldIndustry, fieldListingDate}, ",")
	params := map[string]interface{}{
		"list_status": "L",
	}
	resp, err := fetchTushareRawData("stock_basic", fields, params)
	if err != nil {
		return nil, err
	}
	return parseStockBasicRecord(resp)
}

func parseStockBasicRecord(resp *Response) ([]model.Stock, error) {
	err := resp.anyError()
	if err != nil {
		return nil, err
	}
	var stocks []model.Stock
	for _, item := range resp.Data.Items {
		var stock = model.Stock{}
		for i, field := range resp.Data.Fields {
			switch field {
			case fieldTsCode:
				if str, ok := item[i].(string); ok {
					stock.Code = str
				}
			case fieldName:
				if str, ok := item[i].(string); ok {
					stock.Name = str
				}
			case fieldArea:
				if str, ok := item[i].(string); ok {
					stock.Area = str
				}
			case fieldIndustry:
				if str, ok := item[i].(string); ok {
					stock.Industry = str
				}
			case fieldListingDate:
				if str, ok := item[i].(string); ok {
					stock.ListingDate = str
				}
			default:
				continue
			}
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
