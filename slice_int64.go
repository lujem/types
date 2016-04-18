package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// SliceInt64 is used to handle int64 slice in postgres
type SliceInt64 []int64

// Scan convert the json array ino string slice
func (s *SliceInt64) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return errors.New("unsupported type")
	}
	if len(b) == 0 {
		// this is empty json field. simply create a empty map
		*s = []int64{}
		return nil
	}
	return json.Unmarshal(b, s)
}

// Value try to get the string slice representation in database
func (s SliceInt64) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// String try to get the string slice representation in database
func (s SliceInt64) String() (string, error) {
	res, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
