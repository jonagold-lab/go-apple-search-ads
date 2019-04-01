package searchads

import (
	"encoding/json"
	"fmt"
)

// KeywordDisplayStatus type to represent enum of DisplayStatus (ACTIVE/PAUSED)
type KeywordDisplayStatus byte

// ACTIVE and PAUSED enum values
const (
	KDS_ACTIVE KeywordDisplayStatus = iota
	KDS_PAUSED
)

var (
	_KeywordDisplayStatusNameToValue = map[string]KeywordDisplayStatus{
		"ACTIVE": KDS_ACTIVE,
		"PAUSED": KDS_PAUSED,
	}

	_KeywordDisplayStatusValueToName = map[KeywordDisplayStatus]string{
		KDS_ACTIVE: "ACTIVE",
		KDS_PAUSED: "PAUSED",
	}
)

func init() {
	var v KeywordDisplayStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_KeywordDisplayStatusNameToValue = map[string]KeywordDisplayStatus{
			interface{}(KDS_ACTIVE).(fmt.Stringer).String(): KDS_ACTIVE,
			interface{}(KDS_PAUSED).(fmt.Stringer).String(): KDS_PAUSED,
		}
	}
}

// MarshalJSON is generated so KeywordDisplayStatus satisfies json.Marshaler.
func (r KeywordDisplayStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _KeywordDisplayStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid KeywordDisplayStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so KeywordDisplayStatus satisfies json.Unmarshaler.
func (r *KeywordDisplayStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("KeywordDisplayStatus should be a string, got %s", data)
	}
	v, ok := _KeywordDisplayStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid KeywordDisplayStatus %q", s)
	}
	*r = v
	return nil
}

// ParseKeywordDisplayStatus to turn a String into the KeywordDisplayStatus
func ParseKeywordDisplayStatus(name string) (KeywordDisplayStatus, error) {
	v, ok := _KeywordDisplayStatusNameToValue[name]
	if ok {
		return v, nil
	}
	return KeywordDisplayStatus(0), fmt.Errorf("invalid KeywordDisplayStatus: %s", name)
}

// String to return the String of a KeywordDisplayStatus
func (r KeywordDisplayStatus) String() (string, error) {
	s, ok := _KeywordDisplayStatusValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid KeywordDisplayStatus: %d", r)
	}
	return s, nil
}
