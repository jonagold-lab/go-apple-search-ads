package searchads

import (
	"encoding/json"
	"fmt"
)

// CountryCode type to represent enum of CountryCode (AU/CA)
type CountryCode byte

// AT, AU, BE, CA, CH, CZ, DE, DK, ES, FI, FR, GB, GR, HR, HU, IE, IT, JP, KR, MX, NL, NO, NZ, PT, RO, SE and US enum values
// Anguilla Antigua and Barbuda Argentina Armenia Australia Austria Azerbaijan Bahamas Bahrain Barbados Belarus Belgium Belize Bermuda Bolivia Botswana Brazil British Virgin Islands Brunei Darussalam Bulgaria Burkina Faso Cambodia Canada Cape Verde Cayman Islands China Colombia Costa Rica Croatia Cyprus Czech Republic Denmark Dominica Dominican Republic Ecuador Egypt El Salvador Estonia Fiji Finland France Gambia Germany Gibraltar Greece Grenada Guatemala Guinea‑Bissau Honduras Hong Kong Hungary India Indonesia Ireland Israel Italy Japan Jordan Kazakhstan Kenya Korea, Republic of Kyrgyzstan Kuwait Lao People‘s Democratic Republic Latvia Lebanon Lithuania Luxembourg Macau Malaysia Malta Mayole Mexico Micronesia Moldova Mongolia Mozambique Namibia Nepal Netherlands New Zealand Nicaragua Niger Nigeria Norway Oman Pakistan Panama Papua New Guinea Paraguay Peru Philippines Poland Portugal Qatar Romania Saudi Arabia Singapore Slovakia Slovenia Spain Sri Lanka St. Kitts and Nevis Swaziland Sweden Switzerland Taiwan Tajikistan Thailand Trinidad and Tobago Turkey Turkmenistan Uganda Ukraine United Arab Emirates United Kingdom United States Uzbekistan Venezuela Vietnam Zimbabwe
const (
	AT CountryCode = iota
	AU
	BE
	CA
	CH
	CZ
	DE
	DK
	ES
	FI
	FR
	GB
	GR
	HR
	HU
	IE
	IT
	JP
	KR
	MX
	NL
	NO
	NZ
	PT
	RO
	SE
	US
)

var (
	_CountryCodeNameToValue = map[string]CountryCode{
		"AT": AT,
		"AU": AU,
		"BE": BE,
		"CA": CA,
		"CH": CH,
		"CZ": CZ,
		"DE": DE,
		"DK": DK,
		"ES": ES,
		"FI": FI,
		"FR": FR,
		"GB": GB,
		"GR": GR,
		"HR": HR,
		"HU": HU,
		"IE": IE,
		"IT": IT,
		"JP": JP,
		"KR": KR,
		"MX": MX,
		"NL": NL,
		"NO": NO,
		"NZ": NZ,
		"PT": PT,
		"RO": RO,
		"SE": SE,
		"US": US,
	}

	_CountryCodeValueToName = map[CountryCode]string{
		AT: "AT",
		AU: "AU",
		BE: "BE",
		CA: "CA",
		CH: "CH",
		CZ: "CZ",
		DE: "DE",
		DK: "DK",
		ES: "ES",
		FI: "FI",
		FR: "FR",
		GB: "GB",
		GR: "GR",
		HR: "HR",
		HU: "HU",
		IE: "IE",
		IT: "IT",
		JP: "JP",
		KR: "KR",
		MX: "MX",
		NL: "NL",
		NO: "NO",
		NZ: "NZ",
		PT: "PT",
		RO: "RO",
		SE: "SE",
		US: "US",
	}
)

func init() {
	var v CountryCode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_CountryCodeNameToValue = map[string]CountryCode{
			interface{}(AT).(fmt.Stringer).String(): AT,
			interface{}(AU).(fmt.Stringer).String(): AU,
			interface{}(BE).(fmt.Stringer).String(): BE,
			interface{}(CA).(fmt.Stringer).String(): CA,
			interface{}(CH).(fmt.Stringer).String(): CH,
			interface{}(CZ).(fmt.Stringer).String(): CZ,
			interface{}(DE).(fmt.Stringer).String(): DE,
			interface{}(DK).(fmt.Stringer).String(): DK,
			interface{}(ES).(fmt.Stringer).String(): ES,
			interface{}(FI).(fmt.Stringer).String(): FI,
			interface{}(FR).(fmt.Stringer).String(): FR,
			interface{}(GB).(fmt.Stringer).String(): GB,
			interface{}(GR).(fmt.Stringer).String(): GR,
			interface{}(HR).(fmt.Stringer).String(): HR,
			interface{}(HU).(fmt.Stringer).String(): HU,
			interface{}(IE).(fmt.Stringer).String(): IE,
			interface{}(IT).(fmt.Stringer).String(): IT,
			interface{}(JP).(fmt.Stringer).String(): JP,
			interface{}(KR).(fmt.Stringer).String(): KR,
			interface{}(MX).(fmt.Stringer).String(): MX,
			interface{}(NL).(fmt.Stringer).String(): NL,
			interface{}(NO).(fmt.Stringer).String(): NO,
			interface{}(NZ).(fmt.Stringer).String(): NZ,
			interface{}(PT).(fmt.Stringer).String(): PT,
			interface{}(RO).(fmt.Stringer).String(): RO,
			interface{}(SE).(fmt.Stringer).String(): SE,
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

// String to return the String of a CountryCode
func (r CountryCode) String() (string, error) {
	s, ok := _CountryCodeValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid CountryCode: %d", r)
	}
	return s, nil
}
