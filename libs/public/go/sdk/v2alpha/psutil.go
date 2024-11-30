package sdkv2alphalib

import (
	"github.com/shirou/gopsutil/v4/host"
)

var OSData *OSInformation

type OSInformation struct {
	Device          string `json:"device"`
	Platform        string `json:"platform"`
	Family          string `json:"family"`
	PlatformVersion string `json:"platform_version"`
}

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

func init() {
	GetOSInformation()
}
