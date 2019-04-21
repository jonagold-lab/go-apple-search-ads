package searchads

import (
	"encoding/json"
	"fmt"
)

// Operator type to represent enum of Operator (IN/CONTAINS_ANY/EQUALS/GREATER_THAN/LESS_THAN/STARTSWITH)
type Operator byte

// IN, EQUALS, GREATER_THAN, LESS_THAN and STARTSWITH enum values
const (
	OperatorIn Operator = iota
	OperatorContainsAny
	OperatorEquals
	OperatorGreaterTahan
	OperatorLessThan
	OperatorStartsWith
)

var (
	_OperatorNameToValue = map[string]Operator{
		"IN":           OperatorIn,
		"CONTAINS_ANY": OperatorContainsAny,
		"EQUALS":       OperatorEquals,
		"GREATER_THAN": OperatorGreaterTahan,
		"LESS_THAN":    OperatorLessThan,
		"STARTSWITH":   OperatorStartsWith,
	}

	_OperatorValueToName = map[Operator]string{
		OperatorIn:           "IN",
		OperatorContainsAny:  "CONTAINS_ANY",
		OperatorEquals:       "EQUALS",
		OperatorGreaterTahan: "GREATER_THAN",
		OperatorLessThan:     "LESS_THAN",
		OperatorStartsWith:   "STARTSWITH",
	}
)

func init() {
	var v Operator
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_OperatorNameToValue = map[string]Operator{
			interface{}(OperatorIn).(fmt.Stringer).String():           OperatorIn,
			interface{}(OperatorContainsAny).(fmt.Stringer).String():  OperatorContainsAny,
			interface{}(OperatorEquals).(fmt.Stringer).String():       OperatorEquals,
			interface{}(OperatorGreaterTahan).(fmt.Stringer).String(): OperatorGreaterTahan,
			interface{}(OperatorLessThan).(fmt.Stringer).String():     OperatorLessThan,
			interface{}(OperatorStartsWith).(fmt.Stringer).String():   OperatorStartsWith,
		}
	}
}

// MarshalJSON is generated so Operator satisfies json.Marshaler.
func (r Operator) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _OperatorValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Operator: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Operator satisfies json.Unmarshaler.
func (r *Operator) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Operator should be a string, got %s", data)
	}
	v, ok := _OperatorNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Operator %q", s)
	}
	*r = v
	return nil
}

// ParseOperator to turn a String into the Operator
func ParseOperator(name string) (Operator, error) {
	v, ok := _OperatorNameToValue[name]
	if ok {
		return v, nil
	}
	return Operator(0), fmt.Errorf("invalid Operator: %s", name)
}

// String to return the String of a Operator
func (r Operator) String() (string, error) {
	s, ok := _OperatorValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid Operator: %d", r)
	}
	return s, nil
}
