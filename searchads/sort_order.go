package searchads

import (
	"encoding/json"
	"fmt"
)

// SortOrder type to represent enum of SortOrder (ASCENDING/DESCENDING)
type SortOrder byte

// ASCENDING and DESCENDING enum values
const (
	ASCENDING SortOrder = iota
	DESCENDING
)

var (
	_SortOrderNameToValue = map[string]SortOrder{
		"ASCENDING":  ASCENDING,
		"DESCENDING": DESCENDING,
	}

	_SortOrderValueToName = map[SortOrder]string{
		ASCENDING:  "ASCENDING",
		DESCENDING: "DESCENDING",
	}
)

func init() {
	var v SortOrder
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SortOrderNameToValue = map[string]SortOrder{
			interface{}(ASCENDING).(fmt.Stringer).String():  ASCENDING,
			interface{}(DESCENDING).(fmt.Stringer).String(): DESCENDING,
		}
	}
}

// MarshalJSON is generated so SortOrder satisfies json.Marshaler.
func (r SortOrder) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SortOrderValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid SortOrder: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so SortOrder satisfies json.Unmarshaler.
func (r *SortOrder) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("SortOrder should be a string, got %s", data)
	}
	v, ok := _SortOrderNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid SortOrder %q", s)
	}
	*r = v
	return nil
}

// ParseSortOrder to turn a String into the SortOrder
func ParseSortOrder(name string) (SortOrder, error) {
	v, ok := _SortOrderNameToValue[name]
	if ok {
		return v, nil
	}
	return SortOrder(0), fmt.Errorf("invalid SortOrder: %s", name)
}

// String to return the String of a SortOrder
func (r SortOrder) String() (string, error) {
	s, ok := _SortOrderValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid SortOrder: %d", r)
	}
	return s, nil
}
