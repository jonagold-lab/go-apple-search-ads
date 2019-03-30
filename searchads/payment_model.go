package searchads

import (
	"encoding/json"
	"fmt"
)

// PaymentModel type to represent enum of PaymentModel (LOC/PAYG)
type PaymentModel byte

// LOC and PAYG enum values
const (
	LOC PaymentModel = iota
	PAYG
	EMPTY
)

var (
	_PaymentModelNameToValue = map[string]PaymentModel{
		"LOC":  LOC,
		"PAYG": PAYG,
		"":     EMPTY,
	}

	_PaymentModelValueToName = map[PaymentModel]string{
		LOC:   "LOC",
		PAYG:  "PAYG",
		EMPTY: "",
	}
)

func init() {
	var v PaymentModel
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_PaymentModelNameToValue = map[string]PaymentModel{
			interface{}(LOC).(fmt.Stringer).String():   LOC,
			interface{}(PAYG).(fmt.Stringer).String():  PAYG,
			interface{}(EMPTY).(fmt.Stringer).String(): EMPTY,
		}
	}
}

// MarshalJSON is generated so PaymentModel satisfies json.Marshaler.
func (r PaymentModel) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PaymentModelValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid PaymentModel: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so PaymentModel satisfies json.Unmarshaler.
func (r *PaymentModel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PaymentModel should be a string, got %s", data)
	}
	v, ok := _PaymentModelNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PaymentModel %q", s)
	}
	*r = v
	return nil
}

// ParsePaymentModel to turn a String into the PaymentModel
func ParsePaymentModel(name string) (PaymentModel, error) {
	v, ok := _PaymentModelNameToValue[name]
	if ok {
		return v, nil
	}
	return PaymentModel(0), fmt.Errorf("invalid PaymentModel: %s", name)
}

// String to return the String of a PaymentModel
func (r PaymentModel) String() (string, error) {
	s, ok := _PaymentModelValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid PaymentModel: %d", r)
	}
	return s, nil
}
