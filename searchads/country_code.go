package searchads

import (
	"encoding/json"
	"fmt"
)

// CountryCode type to represent enum of CountryCode
type CountryCode byte

// CountryCode enum values for 46 countries (https://searchads.apple.com/countries-and-regions/)
const (
	// Africa, Middle East, and India
	EG CountryCode = iota
	IN
	IL
	JO
	KW
	LB
	OM
	PK
	QA
	SA
	ZA
	AE
	// Asia Pacific
	AU
	KH
	HK
	ID
	JP
	MO
	MY
	PH
	NZ
	SG
	KR
	TW
	TH
	VN
	// Europe
	AL
	AT
	AZ
	BE
	HR
	CZ
	DK
	FI
	FR
	DE
	GR
	HU
	IE
	IT
	KZ
	NL
	NO
	PL
	PT
	RO
	ES
	SE
	CH
	UA
	GB
	// Latin America and the Caribbean
	AR
	CO
	CL
	EC
	MX
	PE
	// The United States, Canada, and Puerto Rico
	US
	CA
)

var (
	_CountryCodeNameToValue = map[string]CountryCode{
		"AR": AR,
		"AL": AL,
		"IL": IL,
		"AT": AT,
		"AU": AU,
		"AZ": AZ,
		"JO": JO,
		"BE": BE,
		"KW": KW,
		"MO": MO,
		"AE": AE,
		"LB": LB,
		"OM": OM,
		"PK": PK,
		"QA": QA,
		"SA": SA,
		"ZA": ZA,
		"CA": CA,
		"CH": CH,
		"CO": CO,
		"CL": CL,
		"CZ": CZ,
		"DE": DE,
		"DK": DK,
		"EC": EC,
		"EG": EG,
		"ES": ES,
		"FI": FI,
		"FR": FR,
		"GB": GB,
		"GR": GR,
		"HK": HK,
		"HR": HR,
		"HU": HU,
		"IE": IE,
		"IN": IN,
		"ID": ID,
		"IT": IT,
		"JP": JP,
		"KH": KH,
		"KR": KR,
		"KZ": KZ,
		"MX": MX,
		"MY": MY,
		"NL": NL,
		"NO": NO,
		"NZ": NZ,
		"PE": PE,
		"RO": RO,
		"PH": PH,
		"PL": PL,
		"PT": PT,
		"US": US,
		"SE": SE,
		"SG": SG,
		"TW": TW,
		"TH": TH,
		"VN": VN,
		"UA": UA,
	}

	_CountryCodeValueToName = map[CountryCode]string{
		AR: "AR",
		AL: "AL",
		IL: "IL",
		IN: "IN",
		AT: "AT",
		AU: "AU",
		AZ: "AZ",
		JO: "JO",
		BE: "BE",
		KW: "KW",
		MO: "MO",
		AE: "AE",
		LB: "LB",
		OM: "OM",
		PK: "PK",
		QA: "QA",
		SA: "SA",
		ZA: "ZA",
		CA: "CA",
		CH: "CH",
		CO: "CO",
		CL: "CL",
		CZ: "CZ",
		DE: "DE",
		DK: "DK",
		EC: "EC",
		EG: "EG",
		ES: "ES",
		FI: "FI",
		FR: "FR",
		GB: "GB",
		GR: "GR",
		HK: "HK",
		HR: "HR",
		HU: "HU",
		IE: "IE",
		ID: "ID",
		IT: "IT",
		JP: "JP",
		KH: "KH",
		KR: "KR",
		KZ: "KZ",
		MX: "MX",
		MY: "MY",
		NL: "NL",
		NO: "NO",
		NZ: "NZ",
		PE: "PE",
		RO: "RO",
		PH: "PH",
		PL: "PL",
		PT: "PT",
		US: "US",
		SE: "SE",
		SG: "SG",
		TW: "TW",
		TH: "TH",
		VN: "VN",
		UA: "UA",
	}
)

func init() {
	var v CountryCode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_CountryCodeNameToValue = map[string]CountryCode{
			interface{}(AR).(fmt.Stringer).String(): AR,
			interface{}(AL).(fmt.Stringer).String(): AL,
			interface{}(IL).(fmt.Stringer).String(): IL,
			interface{}(IN).(fmt.Stringer).String(): IN,
			interface{}(AT).(fmt.Stringer).String(): AT,
			interface{}(AU).(fmt.Stringer).String(): AU,
			interface{}(AZ).(fmt.Stringer).String(): AZ,
			interface{}(JO).(fmt.Stringer).String(): JO,
			interface{}(BE).(fmt.Stringer).String(): BE,
			interface{}(MO).(fmt.Stringer).String(): MO,
			interface{}(KW).(fmt.Stringer).String(): KW,
			interface{}(AE).(fmt.Stringer).String(): AE,
			interface{}(ZA).(fmt.Stringer).String(): ZA,
			interface{}(LB).(fmt.Stringer).String(): LB,
			interface{}(OM).(fmt.Stringer).String(): OM,
			interface{}(PK).(fmt.Stringer).String(): PK,
			interface{}(QA).(fmt.Stringer).String(): QA,
			interface{}(SA).(fmt.Stringer).String(): SA,
			interface{}(CA).(fmt.Stringer).String(): CA,
			interface{}(CH).(fmt.Stringer).String(): CH,
			interface{}(CO).(fmt.Stringer).String(): CO,
			interface{}(CL).(fmt.Stringer).String(): CL,
			interface{}(CZ).(fmt.Stringer).String(): CZ,
			interface{}(DE).(fmt.Stringer).String(): DE,
			interface{}(DK).(fmt.Stringer).String(): DK,
			interface{}(EC).(fmt.Stringer).String(): EC,
			interface{}(EG).(fmt.Stringer).String(): EG,
			interface{}(ES).(fmt.Stringer).String(): ES,
			interface{}(FI).(fmt.Stringer).String(): FI,
			interface{}(FR).(fmt.Stringer).String(): FR,
			interface{}(GB).(fmt.Stringer).String(): GB,
			interface{}(GR).(fmt.Stringer).String(): GR,
			interface{}(HK).(fmt.Stringer).String(): HK,
			interface{}(HR).(fmt.Stringer).String(): HR,
			interface{}(HU).(fmt.Stringer).String(): HU,
			interface{}(IE).(fmt.Stringer).String(): IE,
			interface{}(ID).(fmt.Stringer).String(): ID,
			interface{}(IT).(fmt.Stringer).String(): IT,
			interface{}(JP).(fmt.Stringer).String(): JP,
			interface{}(KH).(fmt.Stringer).String(): KH,
			interface{}(KR).(fmt.Stringer).String(): KR,
			interface{}(KZ).(fmt.Stringer).String(): KZ,
			interface{}(MX).(fmt.Stringer).String(): MX,
			interface{}(MY).(fmt.Stringer).String(): MY,
			interface{}(NL).(fmt.Stringer).String(): NL,
			interface{}(NO).(fmt.Stringer).String(): NO,
			interface{}(NZ).(fmt.Stringer).String(): NZ,
			interface{}(PE).(fmt.Stringer).String(): PE,
			interface{}(RO).(fmt.Stringer).String(): RO,
			interface{}(PT).(fmt.Stringer).String(): PT,
			interface{}(PH).(fmt.Stringer).String(): PH,
			interface{}(PL).(fmt.Stringer).String(): PL,
			interface{}(US).(fmt.Stringer).String(): US,
			interface{}(SE).(fmt.Stringer).String(): SE,
			interface{}(SG).(fmt.Stringer).String(): SG,
			interface{}(TW).(fmt.Stringer).String(): TW,
			interface{}(TH).(fmt.Stringer).String(): TH,
			interface{}(VN).(fmt.Stringer).String(): VN,
			interface{}(UA).(fmt.Stringer).String(): UA,
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
