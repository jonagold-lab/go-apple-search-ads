package searchads

import (
	"encoding/json"
	"fmt"
)

// KeywordStatus type to represent enum of KeywordStatus (KEYWORD_ACTIVE/KEYWORD_PAUSED)
type KeywordStatus byte

// KEYWORD_ACTIVE and KEYWORD_PAUSED enum values
const (
	KEYWORD_ACTIVE KeywordStatus = iota
	KEYWORD_PAUSED
)

var (
	_KeywordStatusNameToValue = map[string]KeywordStatus{
		"ACTIVE": KEYWORD_ACTIVE,
		"PAUSED": KEYWORD_PAUSED,
	}

	_KeywordStatusValueToName = map[KeywordStatus]string{
		KEYWORD_ACTIVE: "ACTIVE",
		KEYWORD_PAUSED: "PAUSED",
	}
)

func init() {
	var v KeywordStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_KeywordStatusNameToValue = map[string]KeywordStatus{
			interface{}(KEYWORD_ACTIVE).(fmt.Stringer).String(): KEYWORD_ACTIVE,
			interface{}(KEYWORD_PAUSED).(fmt.Stringer).String(): KEYWORD_PAUSED,
		}
	}
}

// MarshalJSON is generated so KeywordStatus satisfies json.Marshaler.
func (r KeywordStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _KeywordStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid KeywordStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so KeywordStatus satisfies json.Unmarshaler.
func (r *KeywordStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("KeywordStatus should be a string, got %s", data)
	}
	v, ok := _KeywordStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid KeywordStatus %q", s)
	}
	*r = v
	return nil
}

// ParseKeywordStatus to turn a String into the KeywordStatus
func ParseKeywordStatus(name string) (KeywordStatus, error) {
	v, ok := _KeywordStatusNameToValue[name]
	if ok {
		return v, nil
	}
	return KeywordStatus(0), fmt.Errorf("invalid KeywordStatus: %s", name)
}

// String to return the String of a KeywordStatus
func (r KeywordStatus) String() (string, error) {
	s, ok := _KeywordStatusValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid KeywordStatus: %d", r)
	}
	return s, nil
}
