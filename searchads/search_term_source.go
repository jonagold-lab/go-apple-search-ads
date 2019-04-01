package searchads

import (
	"encoding/json"
	"fmt"
)

// SearchTermSource type to represent enum of SearchTermSource (AUTO/TARGETED)
type SearchTermSource byte

// AUTO and TARGETED enum values
const (
	AUTO SearchTermSource = iota
	TARGETED
)

var (
	_SearchTermSourceNameToValue = map[string]SearchTermSource{
		"AUTO":     AUTO,
		"TARGETED": TARGETED,
	}

	_SearchTermSourceValueToName = map[SearchTermSource]string{
		AUTO:     "AUTO",
		TARGETED: "TARGETED",
	}
)

func init() {
	var v SearchTermSource
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SearchTermSourceNameToValue = map[string]SearchTermSource{
			interface{}(AUTO).(fmt.Stringer).String():     AUTO,
			interface{}(TARGETED).(fmt.Stringer).String(): TARGETED,
		}
	}
}

// MarshalJSON is generated so SearchTermSource satisfies json.Marshaler.
func (r SearchTermSource) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SearchTermSourceValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid SearchTermSource: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so SearchTermSource satisfies json.Unmarshaler.
func (r *SearchTermSource) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("SearchTermSource should be a string, got %s", data)
	}
	v, ok := _SearchTermSourceNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid SearchTermSource %q", s)
	}
	*r = v
	return nil
}

// ParseSearchTermSource to turn a String into the SearchTermSource
func ParseSearchTermSource(name string) (SearchTermSource, error) {
	v, ok := _SearchTermSourceNameToValue[name]
	if ok {
		return v, nil
	}
	return SearchTermSource(0), fmt.Errorf("invalid SearchTermSource: %s", name)
}

// String to return the String of a SearchTermSource
func (r SearchTermSource) String() (string, error) {
	s, ok := _SearchTermSourceValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid SearchTermSource: %d", r)
	}
	return s, nil
}
