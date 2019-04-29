package searchads

import (
	"encoding/json"
	"fmt"
)

// SearchTermSource type to represent enum of SearchTermSource (AUTO/TARGETED)
type SearchTermSource byte

// AUTO and TARGETED enum values
const (
	SearchTermSourceAuto SearchTermSource = iota
	SearchTermSourceTargeted
)

var (
	_SearchTermSourceNameToValue = map[string]SearchTermSource{
		"AUTO":     SearchTermSourceAuto,
		"TARGETED": SearchTermSourceTargeted,
	}

	_SearchTermSourceValueToName = map[SearchTermSource]string{
		SearchTermSourceAuto:     "AUTO",
		SearchTermSourceTargeted: "TARGETED",
	}
)

func init() {
	var v SearchTermSource
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SearchTermSourceNameToValue = map[string]SearchTermSource{
			interface{}(SearchTermSourceAuto).(fmt.Stringer).String():     SearchTermSourceAuto,
			interface{}(SearchTermSourceTargeted).(fmt.Stringer).String(): SearchTermSourceTargeted,
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
