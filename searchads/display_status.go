package searchads

import (
	"encoding/json"
	"fmt"
)

// DisplayStatus type to represent enum of DisplayStatus (RUNNING/PAUSED)
type DisplayStatus byte

// RUNNING and PAUSED enum values
const (
	DS_RUNNING DisplayStatus = iota
	DS_PAUSED
)

var (
	_DisplayStatusNameToValue = map[string]DisplayStatus{
		"RUNNING": DS_RUNNING,
		"PAUSED":  DS_PAUSED,
	}

	_DisplayStatusValueToName = map[DisplayStatus]string{
		DS_RUNNING: "RUNNING",
		DS_PAUSED:  "PAUSED",
	}
)

func init() {
	var v DisplayStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_DisplayStatusNameToValue = map[string]DisplayStatus{
			interface{}(DS_RUNNING).(fmt.Stringer).String(): DS_RUNNING,
			interface{}(DS_PAUSED).(fmt.Stringer).String():  DS_PAUSED,
		}
	}
}

// MarshalJSON is generated so DisplayStatus satisfies json.Marshaler.
func (r DisplayStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _DisplayStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid DisplayStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so DisplayStatus satisfies json.Unmarshaler.
func (r *DisplayStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("DisplayStatus should be a string, got %s", data)
	}
	v, ok := _DisplayStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid DisplayStatus %q", s)
	}
	*r = v
	return nil
}

// ParseDisplayStatus to turn a String into the DisplayStatus
func ParseDisplayStatus(name string) (DisplayStatus, error) {
	v, ok := _DisplayStatusNameToValue[name]
	if ok {
		return v, nil
	}
	return DisplayStatus(0), fmt.Errorf("invalid DisplayStatus: %s", name)
}

// String to return the String of a DisplayStatus
func (r DisplayStatus) String() (string, error) {
	s, ok := _DisplayStatusValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid DisplayStatus: %d", r)
	}
	return s, nil
}
