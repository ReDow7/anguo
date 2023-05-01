package common

import "time"

func GetLastYearEndDate() string {
	now := time.Now()
	deltaYear := -2
	if now.Month() >= time.May {
		deltaYear = -1
	}
	now = now.AddDate(deltaYear, 0, 0)
	return now.Format("2006") + "1231"
}
