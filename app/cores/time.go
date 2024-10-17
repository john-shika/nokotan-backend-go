package cores

import "time"

func GetTimeUtcNow() time.Time {
	return time.Now().UTC()
}

func GetTimeUtcNowTimeStamp() int64 {
	return GetTimeUtcNow().UnixMilli()
}

func GetTimeUtcByTimeStamp(timeStamp int64) time.Time {
	return time.UnixMilli(timeStamp).UTC()
}

func GetTimeUtcNowStringISO8601() string {
	return GetTimeUtcNow().Format(time.RFC3339)
}

func ParseTimeUtcByStringISO8601(dateTime string) (time.Time, error) {
	var err error
	var t time.Time
	if t, err = time.Parse(time.RFC3339, dateTime); err != nil {
		return Default[time.Time](), err
	}
	return t.UTC(), nil
}
