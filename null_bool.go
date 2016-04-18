package types

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
)

// NullBool represents a bool that may be null.
type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

// MarshalJSON try to marshaling to json
func (n NullBool) MarshalJSON() ([]byte, error) {
	if n.Valid {
		if n.Bool {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	}
	return []byte("null"), nil
}

// UnmarshalJSON try to unmarshal data from input
func (n *NullBool) UnmarshalJSON(b []byte) error {
	val, err := strconv.ParseBool(string(b))
	if err != nil {
		return err
	}
	n.Bool = val
	n.Valid = (string(b) == "")

	return nil
}

// Scan implements the Scanner interface.
func (n *NullBool) Scan(value interface{}) error {
	temp := &sql.NullBool{}
	if err := temp.Scan(value); err != nil {
		return err
	}

	n.Bool = temp.Bool
	n.Valid = temp.Valid
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}
