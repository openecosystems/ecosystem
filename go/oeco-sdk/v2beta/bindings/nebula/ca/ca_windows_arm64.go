//go:build windows && arm64

package nebulav1ca

import "embed"

//go:embed bin/nebula-windows-arm64/nebula-cert.exe
var embeddedFiles embed.FS

var binaryPath = "bin/nebula-windows-arm64/nebula-cert.exe"
