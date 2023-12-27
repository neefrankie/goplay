package chrono

import (
	"fmt"
	"time"
)

// ParseDateTime parses SQL DATE or DATETIME string in specified location.
func ParseDateTime(str string, loc *time.Location) (time.Time, error) {
	base := "0000-00-00 00:00:00.0000000"
	// Zero value
	if str == base[:len(str)] {
		return time.Time{}, nil
	}

	var format string
	switch len(str) {
	case 10, 19: // up to "YYYY-MM-DD HH:MM:SS"
		format = SQLDateTime[:len(str)]
	case 23, 26:
		// Ignore fractional part.
		format = SQLDateTime
	default:
		return time.Time{}, fmt.Errorf("invalid time string: %s", str)
	}

	t, err := time.Parse(format, str)
	if err != nil {
		return time.Time{}, err
	}

	// Adjust location
	if loc != time.UTC {
		y, mo, d := t.Date()
		h, mi, s := t.Clock()
		t, err = time.Date(y, mo, d, h, mi, s, t.Nanosecond(), loc), nil
	}

	return t, err
}
