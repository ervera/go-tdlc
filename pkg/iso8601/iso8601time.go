package iso8601

import (
	"fmt"
	"time"
)

// ITime represents a date time that can be marshaled and unmarshaled following the ISO 8601 convention.
type ITime time.Time

// Now returns the current local ITime.
func NowTime() ITime {
	return ITime(time.Now())
}

// MarshalJSON creates a JSON representation for a date, following the ISO 8601 convention.
func (t ITime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.000-07:00"))
	return []byte(stamp), nil
}

// UnmarshalJSON parses a JSON representation, following the ISO 8601 convention, to a time element.
func (t *ITime) UnmarshalJSON(b []byte) error {
	time, err := time.Parse("\"2006-01-02T15:04:05.000-07:00\"", string(b))
	if err == nil {
		*t = ITime(time)
	}

	return err
}
