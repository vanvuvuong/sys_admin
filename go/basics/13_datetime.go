package basics

import "time"

func Timer() {
	now := time.Now()
	P("Current time:", now)
	P("Current UNIX time:", now.Unix())
	P("Current time with custom format 1:", now.Format("Mon May 2"))
	P("Current time with custom format 2:", now.Format(time.RFC3339)) // like 2019-02-13T13:48:10+01:00
	P("Current time with custom format 3:", now.Format(time.RFC822Z)) // like 13 Feb 19 13:50 +0100
	P("Current time with custom format 4:", now.Format(time.RFC1123)) // like Wed, 13 Feb 2019 13:51:06 CET

	// time in a specific location
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	Check("Error to get time.LoadLocation", err)
	nowHCM := now.In(location)
	P("Current time in HCM:", nowHCM.Format(time.RFC3339Nano))

	// parse time by string with time.RFC3339 format - not tested
	// timeToParse := "2019-02-15 07:33"
	// times, err := time.Parse(time.Now().Format("2019-02-15 07:33"), timeToParse)
	// Check("Error to parse time from string", err)
	// P(timeToParse, " to ", times.Format(time.RFC3339))

	// add time
	now = now.Add(time.Hour * 1)
	P("Next hour: ", now.Format(time.RFC3339))

	// loop overtime
	start, err := time.Parse("2006-01-02", "2019-02-19")
	Check("Error to parse time from string", err)
	end, err := time.Parse("2006-01-02", "2020-07-17")
	Check("Error to parse time from string", err)
	for index := start; index.Unix() < end.Unix(); index = index.AddDate(0, 0, 1) {
		P(index.Format(time.RFC3339))
	}
}
