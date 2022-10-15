package tushare

import (
	"fmt"
	"infra/dal"
	"strings"
	"time"
)

type marketCalendar struct {
	date                string
	open                int
	dateOfLastMarketDay string
}

func (t *marketCalendar) isOpen() bool {
	return t.open == 1
}

func GetTokenFromFile() string {
	token, err := dal.ReadFromFile("token.sec")
	if err != nil {
		panic(fmt.Sprintf("can not read tushare token from file %v", err))
	}
	return token
}

func GetLeastCurrentMarketDate() (string, error) {
	return GetLeastCurrentMarketDateOfGiven("")
}

func GetLeastCurrentMarketDateOfGiven(given string) (string, error) {
	if given == "" {
		given = currentDate()
	}
	tc, err := getCalendarFromTushare(given)
	if err != nil {
		return "", err
	}
	if tc.isOpen() {
		return tc.date, nil
	} else {
		return tc.dateOfLastMarketDay, nil
	}
}

func currentDate() string {
	return time.Now().Format("20060102")
}

func getCalendarFromTushare(callDate string) (*marketCalendar, error) {
	fields := strings.Join([]string{fieldExchange, fieldCalDate, fieldIsOpen, fieldPretradeDate}, ",")
	params := map[string]interface{}{
		fieldExchange: "SSE",
		fieldCalDate:  callDate,
	}
	resp, err := fetchTushareRawData("trade_cal", fields, params)
	if err != nil {
		return nil, err
	}
	return parseToolRecord(resp)
}

func parseToolRecord(resp *Response) (*marketCalendar, error) {
	err := resp.anyError()
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Items) == 0 {
		return nil, fmt.Errorf("encouter an empty data.Items fetched from tushare")
	}
	var ret marketCalendar
	for i, field := range resp.Data.Fields {
		switch field {
		case fieldIsOpen:
			if val, ok := resp.Data.Items[0][i].(int); ok {
				ret.open = val
			}
		case fieldCalDate:
			if str, ok := resp.Data.Items[0][i].(string); ok {
				ret.date = str
			}
		case fieldPretradeDate:
			if str, ok := resp.Data.Items[0][i].(string); ok {
				ret.dateOfLastMarketDay = str
			}
		default:
			continue
		}
	}
	return &ret, nil
}
