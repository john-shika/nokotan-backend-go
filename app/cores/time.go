package cores

import (
	"time"
)

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

type TimeAnyImpl interface {
	time.Time | string | int64 | int
}

func GetTimeUtcFromTimeAny(value any) time.Time {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return Default[time.Time]()
	}
	value = val.Interface()
	switch value.(type) {
	case time.Time:
		return value.(time.Time).UTC()
	case string:
		return Unwrap(ParseTimeUtcByStringISO8601(value.(string)))
	case int64:
		return GetTimeUtcByTimeStamp(value.(int64))
	case int:
		return GetTimeUtcByTimeStamp(int64(value.(int)))
	default:
		return Default[time.Time]()
	}
}

func GetTimeUtcFromTimeAnyStrict[V TimeAnyImpl](value V) time.Time {
	return GetTimeUtcFromTimeAny(value)
}
