package searchads

import (
	"encoding/json"
	"fmt"
)

// MatchType type to represent enum of MatchType (EXACT/BORAD)
type MatchType byte

// EXACT and BORAD enum values
const (
	EXACT MatchType = iota
	BORAD
)

var (
	_MatchTypeNameToValue = map[string]MatchType{
		"EXACT": EXACT,
		"BORAD": BORAD,
	}

	_MatchTypeValueToName = map[MatchType]string{
		EXACT: "EXACT",
		BORAD: "BORAD",
	}
)

func init() {
	var v MatchType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MatchTypeNameToValue = map[string]MatchType{
			interface{}(EXACT).(fmt.Stringer).String(): EXACT,
			interface{}(BORAD).(fmt.Stringer).String(): BORAD,
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
