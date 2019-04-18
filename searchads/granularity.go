package searchads

import (
	"encoding/json"
	"fmt"
)

// Granularity type to represent enum of Granularity (HOURLY/DAILY/WEEKLY/MONTHLY)
type Granularity byte

// HOURLY, DAILY, WEEKLY and MONTHLY enum values
const (
	HOURLY Granularity = iota
	DAILY
	WEEKLY
	MONTHLY
)

var (
	_GranularityNameToValue = map[string]Granularity{
		"HOURLY":  HOURLY,
		"DAILY":   DAILY,
		"WEEKLY":  WEEKLY,
		"MONTHLY": MONTHLY,
	}

	_GranularityValueToName = map[Granularity]string{
		HOURLY:  "HOURLY",
		DAILY:   "DAILY",
		WEEKLY:  "WEEKLY",
		MONTHLY: "MONTHLY",
	}
)

func init() {
	var v Granularity
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_GranularityNameToValue = map[string]Granularity{
			interface{}(HOURLY).(fmt.Stringer).String():  HOURLY,
			interface{}(DAILY).(fmt.Stringer).String():   DAILY,
			interface{}(WEEKLY).(fmt.Stringer).String():  WEEKLY,
			interface{}(MONTHLY).(fmt.Stringer).String(): MONTHLY,
		}
	}
}

// MarshalJSON is generated so Granularity satisfies json.Marshaler.
func (r Granularity) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _GranularityValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Granularity: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Granularity satisfies json.Unmarshaler.
func (r *Granularity) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Granularity should be a string, got %s", data)
	}
	v, ok := _GranularityNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Granularity %q", s)
	}
	*r = v
	return nil
}

// ParseGranularity to turn a String into the Granularity
func ParseGranularity(name string) (Granularity, error) {
	v, ok := _GranularityNameToValue[name]
	if ok {
		return v, nil
	}
	return Granularity(0), fmt.Errorf("invalid Granularity: %s", name)
}

// String to return the String of a Granularity
func (r Granularity) String() (string, error) {
	s, ok := _GranularityValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid Granularity: %d", r)
	}
	return s, nil
}
