package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// GenericJSONField is used to handle generic json data in postgres
type GenericJSONField map[string]interface{}

// Scan convert the json array ino string slice
func (f *GenericJSONField) Scan(src interface{}) error {
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
		*f = make(GenericJSONField)
		return nil
	}
	return json.Unmarshal(b, f)
}

// Value try to get the string slice representation in database
func (f GenericJSONField) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// String try to get the string slice representation in database
func (f GenericJSONField) String() (string, error) {
	res, err := json.Marshal(f)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
