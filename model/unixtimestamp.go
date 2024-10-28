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

func (v UnixTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Time.Unix())
}

func (v *UnixTimestamp) UnmarshalJSON(src []byte) error {
	value, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}
	v.Time = time.Unix(value, 0)
	return nil
}
