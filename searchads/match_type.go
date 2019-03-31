package searchads

import (
	"encoding/json"
	"fmt"
)

// MatchType type to represent enum of MatchType (EXACT/BROAD)
type MatchType byte

// EXACT and BROAD enum values
const (
	EXACT MatchType = iota
	BROAD
)

var (
	_MatchTypeNameToValue = map[string]MatchType{
		"EXACT": EXACT,
		"BROAD": BROAD,
	}

	_MatchTypeValueToName = map[MatchType]string{
		EXACT: "EXACT",
		BROAD: "BROAD",
	}
)

func init() {
	var v MatchType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MatchTypeNameToValue = map[string]MatchType{
			interface{}(EXACT).(fmt.Stringer).String(): EXACT,
			interface{}(BROAD).(fmt.Stringer).String(): BROAD,
		}
	}
}

// MarshalJSON is generated so MatchType satisfies json.Marshaler.
func (r MatchType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MatchTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid MatchType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so MatchType satisfies json.Unmarshaler.
func (r *MatchType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MatchType should be a string, got %s", data)
	}
	v, ok := _MatchTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid MatchType %q", s)
	}
	*r = v
	return nil
}

// ParseMatchType to turn a String into the MatchType
func ParseMatchType(name string) (MatchType, error) {
	v, ok := _MatchTypeNameToValue[name]
	if ok {
		return v, nil
	}
	return MatchType(0), fmt.Errorf("invalid MatchType: %s", name)
}

// String to return the String of a MatchType
func (r MatchType) String() (string, error) {
	s, ok := _MatchTypeValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid MatchType: %d", r)
	}
	return s, nil
}
