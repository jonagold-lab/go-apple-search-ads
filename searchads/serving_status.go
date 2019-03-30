package searchads

import (
	"encoding/json"
	"fmt"
)

// ServingStatus type to represent enum of ServingStatus (RUNNING/NOT_RUNNING)
type ServingStatus byte

// RUNNING and NOT_RUNNING enum values
const (
	RUNNING ServingStatus = iota
	NOT_RUNNING
)

var (
	_ServingStatusNameToValue = map[string]ServingStatus{
		"RUNNING":     RUNNING,
		"NOT_RUNNING": NOT_RUNNING,
	}

	_ServingStatusValueToName = map[ServingStatus]string{
		RUNNING:     "RUNNING",
		NOT_RUNNING: "NOT_RUNNING",
	}
)

func init() {
	var v ServingStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ServingStatusNameToValue = map[string]ServingStatus{
			interface{}(RUNNING).(fmt.Stringer).String():     RUNNING,
			interface{}(NOT_RUNNING).(fmt.Stringer).String(): NOT_RUNNING,
		}
	}
}

// MarshalJSON is generated so ServingStatus satisfies json.Marshaler.
func (r ServingStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ServingStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ServingStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ServingStatus satisfies json.Unmarshaler.
func (r *ServingStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ServingStatus should be a string, got %s", data)
	}
	v, ok := _ServingStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ServingStatus %q", s)
	}
	*r = v
	return nil
}

// ParseServingStatus to turn a String into the ServingStatus
func ParseServingStatus(name string) (ServingStatus, error) {
	v, ok := _ServingStatusNameToValue[name]
	if ok {
		return v, nil
	}
	return ServingStatus(0), fmt.Errorf("invalid ServingStatus: %s", name)
}

// String to return the String of a ServingStatus
func (r ServingStatus) String() (string, error) {
	s, ok := _ServingStatusValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid ServingStatus: %d", r)
	}
	return s, nil
}
