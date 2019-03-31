package serachads

type TargetingDimensions struct {
	Age            Age            `json:"age,omitempty"`
	Gender         Gender         `json:"gende,omitemptyr"`
	Country        Country        `json:"country,omitempty"`
	AdminArea      AdminArea      `json:"adminArea,omitempty"`
	Locality       Locality       `json:"locality,omitempty"`
	DeviceClass    DeviceClass    `json:"deviceClass,omitempty"`
	Daypart        Daypart        `json:"daypart,omitempty"`
	AppDownloaders AppDownloaders `json:"appDownloaders,omitempty"`
}

type Country struct {
	Included []string `json:"included,omitempty"`
}

type Age struct {
	Included []AgeObj `json:"included,omitempty"`
}

type AgeObj struct {
	MinAge string `json:"minAge,omitempty"`
	MaxAge string `json:"maxAge,omitempty"`
}

type Gender struct {
	Included []string `json:"included,omitempty"`
}

type AdminArea struct {
	Included []string `json:"included,omitempty"`
}

type Locality struct {
	Included []string `json:"included,omitempty"`
}

type Daypart struct {
	UserTime UserTime `json:"userTime,omitempty"`
}

type UserTime struct {
	Included []string `json:"included,omitempty"`
}

type DeviceClass struct {
	Included []string `json:"included,omitempty"`
}
type AppDownloaders struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}
