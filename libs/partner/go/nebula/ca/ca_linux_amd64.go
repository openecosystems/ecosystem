//go:build linux && amd64

package nebulav1ca

import "embed"

//go:embed bin/nebula-linux-amd64/nebula-cert
var embeddedFiles embed.FS

var binaryPath = "bin/nebula-linux-amd64/nebula-cert"
