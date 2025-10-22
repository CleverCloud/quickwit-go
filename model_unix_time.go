package quickwit

import (
	"encoding/json"
	"time"
)

// UnixTime is our magic type
type UnixTime struct {
	time.Time
}

// UnmarshalJSON is the method that satisfies the Unmarshaller interface
func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}
