package searchads

import (
	"encoding/json"
	"fmt"
)

// Operator type to represent enum of Operator (IN/EQUALS)
type Operator byte

// IN, EQUALS, GREATER_THAN, LESS_THAN and STARTSWITH enum values
const (
	IN Operator = iota
	EQUALS
	GREATER_THAN
	LESS_THAN
	STARTSWITH
)

var (
	_OperatorNameToValue = map[string]Operator{
		"IN":           IN,
		"EQUALS":       EQUALS,
		"GREATER_THAN": GREATER_THAN,
		"LESS_THAN":    LESS_THAN,
		"STARTSWITH":   STARTSWITH,
	}

	_OperatorValueToName = map[Operator]string{
		IN:           "IN",
		EQUALS:       "EQUALS",
		GREATER_THAN: "GREATER_THAN",
		LESS_THAN:    "LESS_THAN",
		STARTSWITH:   "STARTSWITH",
	}
)

func init() {
	var v Operator
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_OperatorNameToValue = map[string]Operator{
			interface{}(IN).(fmt.Stringer).String():           IN,
			interface{}(EQUALS).(fmt.Stringer).String():       EQUALS,
			interface{}(GREATER_THAN).(fmt.Stringer).String(): GREATER_THAN,
			interface{}(LESS_THAN).(fmt.Stringer).String():    LESS_THAN,
			interface{}(STARTSWITH).(fmt.Stringer).String():   STARTSWITH,
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
