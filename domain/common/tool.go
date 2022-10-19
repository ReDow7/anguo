package common

import "time"

func GetLastYearEndDate() string {
	now := time.Now()
	now = now.AddDate(-1, 0, 0)
	return now.Format("2006") + "1231"
}
