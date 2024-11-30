//go:build darwin && arm64

package nebulav1ca

import "embed"

//go:embed bin/nebula-darwin/nebula-cert
var embeddedFiles embed.FS

var binaryPath = "bin/nebula-darwin/nebula-cert"
