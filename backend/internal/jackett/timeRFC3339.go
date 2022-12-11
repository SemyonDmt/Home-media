package jackett

import (
	"fmt"
	"strings"
	"time"
)

type TimeRFC3339 struct {
	time.Time
}

func (ct *TimeRFC3339) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(time.RFC3339, s)

	ct.Time = t
	if err != nil {
		fmt.Println("Parse", err)
	}
	return nil
}
