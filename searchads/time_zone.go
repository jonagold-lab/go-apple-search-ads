package searchads

import (
	"encoding/json"
	"fmt"
)

// TimeZone type to represent enum of TimeZone (ORTZ/UTC)
type TimeZone byte

// ORTZ and UTC enum values
const (
	ORTZ TimeZone = iota
	UTC
)

var (
	_TimeZoneNameToValue = map[string]TimeZone{
		"ORTZ": ORTZ,
		"UTC":  UTC,
	}

	_TimeZoneValueToName = map[TimeZone]string{
		ORTZ: "ORTZ",
		UTC:  "UTC",
	}
)

func init() {
	var v TimeZone
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TimeZoneNameToValue = map[string]TimeZone{
			interface{}(ORTZ).(fmt.Stringer).String(): ORTZ,
			interface{}(UTC).(fmt.Stringer).String():  UTC,
		}
	}
}

// MarshalJSON is generated so TimeZone satisfies json.Marshaler.
func (r TimeZone) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TimeZoneValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid TimeZone: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TimeZone satisfies json.Unmarshaler.
func (r *TimeZone) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TimeZone should be a string, got %s", data)
	}
	v, ok := _TimeZoneNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid TimeZone %q", s)
	}
	*r = v
	return nil
}

// ParseTimeZone to turn a String into the TimeZone
func ParseTimeZone(name string) (TimeZone, error) {
	v, ok := _TimeZoneNameToValue[name]
	if ok {
		return v, nil
	}
	return TimeZone(0), fmt.Errorf("invalid TimeZone: %s", name)
}

// String to return the String of a TimeZone
func (r TimeZone) String() (string, error) {
	s, ok := _TimeZoneValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid TimeZone: %d", r)
	}
	return s, nil
}
