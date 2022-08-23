package iso8601

import (
	"fmt"
	"time"
)

// ISO8601Date represents a date time that can be marshaled and unmarshaled following the ISO 8601 convention.
type IDate time.Time

// Now returns the current local ISO8601Date.
func NowDate() IDate {
	return IDate(time.Now())
}

// MarshalJSON creates a JSON representation for a date, following the ISO 8601 convention.
func (t IDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

// UnmarshalJSON parses a JSON representation, following the ISO 8601 convention, to a time element.
func (t *IDate) UnmarshalJSON(b []byte) error {
	time, err := time.Parse("\"2006-01-02\"", string(b))
	if err == nil {
		*t = IDate(time)
	}

	return err
}
