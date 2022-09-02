package iso8601

import (
	"fmt"
	"time"
)

type IDate time.Time

func NowDate() IDate {
	return IDate(time.Now())
}

func (t IDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (t *IDate) UnmarshalJSON(b []byte) error {
	time, err := time.Parse("\"2006-01-02\"", string(b))
	if err == nil {
		*t = IDate(time)
	}

	return err
}
