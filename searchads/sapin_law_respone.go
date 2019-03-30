package searchads

import (
	"encoding/json"
	"fmt"
)

// SapinLawResponse type to represent enum of SapinLawResponse (NOT_ANSWERED/FRENCH_BUSINESS)
type SapinLawResponse byte

// NOT_ANSWERED, FRENCH_BUSINESS and NOT_FRENCH_BUSINESS enum values
const (
	NOT_ANSWERED SapinLawResponse = iota
	FRENCH_BUSINESS
	NOT_FRENCH_BUSINESS
)

var (
	_SapinLawResponseNameToValue = map[string]SapinLawResponse{
		"NOT_ANSWERED":        NOT_ANSWERED,
		"FRENCH_BUSINESS":     FRENCH_BUSINESS,
		"NOT_FRENCH_BUSINESS": NOT_FRENCH_BUSINESS,
	}

	_SapinLawResponseValueToName = map[SapinLawResponse]string{
		NOT_ANSWERED:        "NOT_ANSWERED",
		FRENCH_BUSINESS:     "FRENCH_BUSINESS",
		NOT_FRENCH_BUSINESS: "NOT_FRENCH_BUSINESS",
	}
)

func init() {
	var v SapinLawResponse
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SapinLawResponseNameToValue = map[string]SapinLawResponse{
			interface{}(NOT_ANSWERED).(fmt.Stringer).String():        NOT_ANSWERED,
			interface{}(FRENCH_BUSINESS).(fmt.Stringer).String():     FRENCH_BUSINESS,
			interface{}(NOT_FRENCH_BUSINESS).(fmt.Stringer).String(): NOT_FRENCH_BUSINESS,
		}
	}
}

// MarshalJSON is generated so SapinLawResponse satisfies json.Marshaler.
func (r SapinLawResponse) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SapinLawResponseValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid SapinLawResponse: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so SapinLawResponse satisfies json.Unmarshaler.
func (r *SapinLawResponse) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("SapinLawResponse should be a string, got %s", data)
	}
	v, ok := _SapinLawResponseNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid SapinLawResponse %q", s)
	}
	*r = v
	return nil
}

// ParseSapinLawResponse to turn a String into the SapinLawResponse
func ParseSapinLawResponse(name string) (SapinLawResponse, error) {
	v, ok := _SapinLawResponseNameToValue[name]
	if ok {
		return v, nil
	}
	return SapinLawResponse(0), fmt.Errorf("invalid SapinLawResponse: %s", name)
}

// String to return the String of a SapinLawResponse
func (r SapinLawResponse) String() (string, error) {
	s, ok := _SapinLawResponseValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid SapinLawResponse: %d", r)
	}
	return s, nil
}
