package timeutil

import "time"

const (
	FormatTimeLong  = "Mon, 02 Jan 2006 15:04"
	FormatTimeShort = "02 Jan 2006 3:04PM"
)

func ParseTimeByTimeZone(t time.Time, timezone string) time.Time {
	if timezone != "" {
		loc, _ := time.LoadLocation(timezone)
		timeFollowTimezone := time.Date(t.In(time.UTC).Year(), t.In(time.UTC).Month(), t.In(time.UTC).Day(), t.In(time.UTC).Hour(), t.In(time.UTC).Minute(), 0, 0, loc).Unix()
		return time.Unix(timeFollowTimezone, 0)
	}
	return t
}
