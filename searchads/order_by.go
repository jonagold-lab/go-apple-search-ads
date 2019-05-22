package searchads

import (
	"encoding/json"
	"fmt"
)

// OrderBy type to represent enum of OrderBy
type OrderBy byte

const (
	OrderByAdminArea OrderBy = iota
	OrderByAgeRange
	OrderByCountryCode
	OrderByCountryOrRegion
	OrderByDeviceClass
	OrderByGender
	OrderByLocality
	OrderByAdGroupID
	OrderByID
	OrderByImpressions
)

var (
	_OrderByNameToValue = map[string]OrderBy{
		"adminArea":       OrderByAdminArea,
		"ageRange":        OrderByAgeRange,
		"countryCode":     OrderByCountryCode,
		"countryOrRegion": OrderByCountryOrRegion,
		"deviceClass":     OrderByDeviceClass,
		"gender":          OrderByGender,
		"locality":        OrderByLocality,
		"adGroupId":       OrderByAdGroupID,
		"id":              OrderByID,
		"impressions":     OrderByImpressions,
	}

	_OrderByValueToName = map[OrderBy]string{
		OrderByAdminArea:       "adminArea",
		OrderByAgeRange:        "ageRange",
		OrderByCountryCode:     "countryCode",
		OrderByCountryOrRegion: "countryOrRegion",
		OrderByDeviceClass:     "deviceClass",
		OrderByGender:          "gender",
		OrderByLocality:        "locality",
		OrderByAdGroupID:       "adGroupId",
		OrderByID:              "id",
		OrderByImpressions:     "impressions",
	}
)

func init() {
	var v OrderBy
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_OrderByNameToValue = map[string]OrderBy{
			interface{}(OrderByAdminArea).(fmt.Stringer).String():       OrderByAdminArea,
			interface{}(OrderByAgeRange).(fmt.Stringer).String():        OrderByAgeRange,
			interface{}(OrderByCountryCode).(fmt.Stringer).String():     OrderByCountryCode,
			interface{}(OrderByCountryOrRegion).(fmt.Stringer).String(): OrderByCountryOrRegion,
			interface{}(OrderByDeviceClass).(fmt.Stringer).String():     OrderByDeviceClass,
			interface{}(OrderByGender).(fmt.Stringer).String():          OrderByGender,
			interface{}(OrderByLocality).(fmt.Stringer).String():        OrderByLocality,
			interface{}(OrderByAdGroupID).(fmt.Stringer).String():       OrderByAdGroupID,
			interface{}(OrderByID).(fmt.Stringer).String():              OrderByID,
			interface{}(OrderByImpressions).(fmt.Stringer).String():     OrderByImpressions,
		}
	}
}

// MarshalJSON is generated so OrderBy satisfies json.Marshaler.
func (r OrderBy) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _OrderByValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid OrderBy: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so OrderBy satisfies json.Unmarshaler.
func (r *OrderBy) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("OrderBy should be a string, got %s", data)
	}
	v, ok := _OrderByNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid OrderBy %q", s)
	}
	*r = v
	return nil
}

// ParseOrderBy to turn a String into the OrderBy
func ParseOrderBy(name string) (OrderBy, error) {
	v, ok := _OrderByNameToValue[name]
	if ok {
		return v, nil
	}
	return OrderBy(0), fmt.Errorf("invalid OrderBy: %s", name)
}

// String to return the String of a OrderBy
func (r OrderBy) String() (string, error) {
	s, ok := _OrderByValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid OrderBy: %d", r)
	}
	return s, nil
}
