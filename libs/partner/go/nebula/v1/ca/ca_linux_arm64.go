//go:build linux && arm64

package nebulav1ca

import "embed"

//go:embed bin/nebula-linux-arm64/nebula-cert
var embeddedFiles embed.FS

var binaryPath = "bin/nebula-linux-arm64/nebula-cert"
