package searchads

import (
	"encoding/json"
	"fmt"
)

// CountryCode type to represent enum of CountryCode (AU/CA)
type CountryCode byte

// AU, CA, CH, DE, ES, FR, GB, IT, JP, KR, MX, NZ and US enum values
const (
	AU CountryCode = iota
	CA
	CH
	DE
	ES
	FR
	GB
	IT
	JP
	KR
	MX
	NZ
	US
)

var (
	_CountryCodeNameToValue = map[string]CountryCode{
		"AU": AU,
		"CA": CA,
		"CH": CH,
		"DE": DE,
		"ES": ES,
		"FR": FR,
		"GB": GB,
		"IT": IT,
		"JP": JP,
		"KR": KR,
		"MX": MX,
		"NZ": NZ,
		"US": US,
	}

	_CountryCodeValueToName = map[CountryCode]string{
		AU: "AU",
		CA: "CA",
		CH: "CH",
		DE: "DE",
		ES: "ES",
		FR: "FR",
		GB: "GB",
		IT: "IT",
		JP: "JP",
		KR: "KR",
		MX: "MX",
		NZ: "NZ",
		US: "US",
	}
)

func init() {
	var v CountryCode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_CountryCodeNameToValue = map[string]CountryCode{
			interface{}(AU).(fmt.Stringer).String(): AU,
			interface{}(CA).(fmt.Stringer).String(): CA,
			interface{}(CH).(fmt.Stringer).String(): CH,
			interface{}(DE).(fmt.Stringer).String(): DE,
			interface{}(ES).(fmt.Stringer).String(): ES,
			interface{}(FR).(fmt.Stringer).String(): FR,
			interface{}(GB).(fmt.Stringer).String(): GB,
			interface{}(IT).(fmt.Stringer).String(): IT,
			interface{}(JP).(fmt.Stringer).String(): JP,
			interface{}(KR).(fmt.Stringer).String(): KR,
			interface{}(MX).(fmt.Stringer).String(): MX,
			interface{}(NZ).(fmt.Stringer).String(): NZ,
			interface{}(US).(fmt.Stringer).String(): US,
		}
	}
}

// MarshalJSON is generated so CountryCode satisfies json.Marshaler.
func (r CountryCode) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _CountryCodeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid CountryCode: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so CountryCode satisfies json.Unmarshaler.
func (r *CountryCode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CountryCode should be a string, got %s", data)
	}
	v, ok := _CountryCodeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid CountryCode %q", s)
	}
	*r = v
	return nil
}

// ParseCountryCode to turn a String into the CountryCode
func ParseCountryCode(name string) (CountryCode, error) {
	v, ok := _CountryCodeNameToValue[name]
	if ok {
		return v, nil
	}
	return CountryCode(0), fmt.Errorf("invalid CountryCode: %s", name)
}
