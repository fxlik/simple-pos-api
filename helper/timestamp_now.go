package helper

import "time"

func GetTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)
	return now
}
