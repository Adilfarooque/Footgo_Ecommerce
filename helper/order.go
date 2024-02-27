package helper

import "time"

func GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {
	//Get the current time
	endDate := time.Now()
	//Calculate start and end dates based on the time period
	if timePeriod == "day" {
		startDate := endDate.AddDate(0, 0, -1) //Subtract 1 day from the current time
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, 0, -6) //Subtract 6 day from the current time
		return startDate, endDate
	}

	if timePeriod == "year" {
		startTime := endDate.AddDate(-1, 0, 0) //Subtract 1 year from the current time
		return startTime, endDate
	}
	//Default case: return a time period of 7 days(1week)
	return endDate.AddDate(0, 0, -6), endDate
}
