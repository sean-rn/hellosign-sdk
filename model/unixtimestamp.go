package model

import (
	"encoding/json"
	"strconv"
	"time"
)

// UnixTimestamp models a timestamp encoded in JSON as an integer: Unix Timestamp in seconds.
type UnixTimestamp struct {
	time.Time
}

// MarshalJSON encodes into JSON
func (v UnixTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Time.Unix())
}

// UnmarshalJSON parses it from JSON.  Some parts of the API encode as a number,
// but other parts encode as a numeric string, so handle both.
func (v *UnixTimestamp) UnmarshalJSON(src []byte) error {
	if len(src) > len(`""`) && src[0] == '"' && src[len(src)-1] == '"' {
		src = src[1 : len(src)-1]
	}

	value, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}
	v.Time = time.Unix(value, 0)
	return nil
}
