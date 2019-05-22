package searchads

import (
	"encoding/json"
	"fmt"
)

// GroupBy type to represent enum of GroupBy
type GroupBy byte

const (
	GroupByAdminArea GroupBy = iota
	GroupByAgeRange
	GroupByCountryCode
	GroupByCountryOrRegion
	GroupByDeviceClass
	GroupByGender
	GroupByLocality
	GroupByAdGroupID
)

var (
	_GroupByNameToValue = map[string]GroupBy{
		"adminArea":       GroupByAdminArea,
		"ageRange":        GroupByAgeRange,
		"countryCode":     GroupByCountryCode,
		"countryOrRegion": GroupByCountryOrRegion,
		"deviceClass":     GroupByDeviceClass,
		"gender":          GroupByGender,
		"locality":        GroupByLocality,
		"adGroupId":       GroupByAdGroupID,
	}

	_GroupByValueToName = map[GroupBy]string{
		GroupByAdminArea:       "adminArea",
		GroupByAgeRange:        "ageRange",
		GroupByCountryCode:     "countryCode",
		GroupByCountryOrRegion: "countryOrRegion",
		GroupByDeviceClass:     "deviceClass",
		GroupByGender:          "gender",
		GroupByLocality:        "locality",
		GroupByAdGroupID:       "adGroupId",
	}
)

func init() {
	var v GroupBy
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_GroupByNameToValue = map[string]GroupBy{
			interface{}(GroupByAdminArea).(fmt.Stringer).String():       GroupByAdminArea,
			interface{}(GroupByAgeRange).(fmt.Stringer).String():        GroupByAgeRange,
			interface{}(GroupByCountryCode).(fmt.Stringer).String():     GroupByCountryCode,
			interface{}(GroupByCountryOrRegion).(fmt.Stringer).String(): GroupByCountryOrRegion,
			interface{}(GroupByDeviceClass).(fmt.Stringer).String():     GroupByDeviceClass,
			interface{}(GroupByGender).(fmt.Stringer).String():          GroupByGender,
			interface{}(GroupByLocality).(fmt.Stringer).String():        GroupByLocality,
			interface{}(GroupByAdGroupID).(fmt.Stringer).String():       GroupByAdGroupID,
		}
	}
}

// MarshalJSON is generated so GroupBy satisfies json.Marshaler.
func (r GroupBy) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _GroupByValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid GroupBy: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so GroupBy satisfies json.Unmarshaler.
func (r *GroupBy) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("GroupBy should be a string, got %s", data)
	}
	v, ok := _GroupByNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid GroupBy %q", s)
	}
	*r = v
	return nil
}

// ParseGroupBy to turn a String into the GroupBy
func ParseGroupBy(name string) (GroupBy, error) {
	v, ok := _GroupByNameToValue[name]
	if ok {
		return v, nil
	}
	return GroupBy(0), fmt.Errorf("invalid GroupBy: %s", name)
}

// String to return the String of a GroupBy
func (r GroupBy) String() (string, error) {
	s, ok := _GroupByValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid GroupBy: %d", r)
	}
	return s, nil
}
