package iso8601

import (
	"fmt"
	"time"
)

type ITime time.Time

func NowTime() ITime {
	return ITime(time.Now())
}

func (t ITime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.000-07:00"))
	return []byte(stamp), nil
}

func (t *ITime) UnmarshalJSON(b []byte) error {
	time, err := time.Parse("\"2006-01-02T15:04:05.000-07:00\"", string(b))
	if err == nil {
		*t = ITime(time)
	}

	return err
}
