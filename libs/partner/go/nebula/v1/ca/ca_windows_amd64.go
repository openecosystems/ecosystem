//go:build windows && amd64

package nebulav1ca

import "embed"

//go:embed bin/nebula-windows-amd64/nebula-cert
var embeddedFiles embed.FS

var binaryPath = "bin/nebula-windows-amd64/nebula-cert"
