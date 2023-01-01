package jackett

import (
	"strings"
	"time"
)

type TimeRFC3339 struct {
	time.Time
}

func (ct *TimeRFC3339) UnmarshalJSON(b []byte) (err error) {
	str := strings.Trim(string(b), `"`)
	if str == "0001-01-01T00:00:00" {
	} else if len(str) == 19 {
		ct.Time, err = time.Parse(time.RFC3339, str+"Z")
	} else {
		ct.Time, err = time.Parse(time.RFC3339, str)
	}
	return
}
