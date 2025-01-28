package sdkv2alphalib

import (
	"github.com/shirou/gopsutil/v4/host"
)

// OSData holds information about the operating system including platform, version, family, and device details.
var OSData *OSInformation

// OSInformation represents details about the operating system, including device type, platform, family, and version.
type OSInformation struct {
	Device          string `json:"device"`
	Platform        string `json:"platform"`
	Family          string `json:"family"`
	PlatformVersion string `json:"platform_version"`
}

// GetOSInformation retrieves the operating system information including platform, family, and platform version.
// Returns a pointer to an OSInformation struct or nil if an error occurs.
func GetOSInformation() *OSInformation {
	platform, family, pver, err := host.PlatformInformation()
	if err != nil {
		return nil
	}

	OSData = &OSInformation{
		Device:          "",
		Platform:        platform,
		Family:          family,
		PlatformVersion: pver,
	}

	return OSData
}

// init initializes the OSData variable by fetching operating system information using the GetOSInformation function.
func init() {
	GetOSInformation()
}
