package sdkv2betalib

import (
	"regexp"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
)

// OSData holds information about the operating system including platform, version, family, and device details.
var OSData *OSInformation

// OSInformation represents details about the operating system, including device type, platform, family, and version.
type OSInformation struct {
	Device            string `json:"device"`
	Platform          string `json:"platform"`
	Family            string `json:"family"`
	PlatformVersion   string `json:"platform_version"`
	Hostname          string `json:"hostname"`
	SanitizedHostname string `json:"sanitized_hostname"`
}

// GetOSInformation retrieves the operating system information including platform, family, and platform version.
// Returns a pointer to an OSInformation struct or nil if an error occurs.
// TODO: return an error here
func GetOSInformation() *OSInformation {
	platform, family, pver, err := host.PlatformInformation()
	if err != nil {
		return nil
	}

	info, err := host.Info()
	if err != nil {
		return nil
	}

	OSData = &OSInformation{
		Device:            "",
		Platform:          platform,
		Family:            family,
		PlatformVersion:   pver,
		Hostname:          info.Hostname,
		SanitizedHostname: sanitizeHostname(info.Hostname),
	}

	return OSData
}

// init initializes the OSData variable by fetching operating system information using the GetOSInformation function.
func init() {
	GetOSInformation()
}

func sanitizeHostname(hostname string) string {
	// Convert to lowercase (DNS is case-insensitive)
	hostname = strings.ToLower(hostname)

	// Remove invalid characters (allow only letters, digits, and hyphens)
	re := regexp.MustCompile(`[^a-z0-9-]`)
	hostname = re.ReplaceAllString(hostname, "")

	// Trim leading and trailing hyphens
	hostname = strings.Trim(hostname, "-")

	// Enforce max length of 63 characters per segment
	if len(hostname) > 63 {
		hostname = hostname[:63]
	}

	return hostname
}
