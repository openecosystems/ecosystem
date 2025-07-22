package sdkv2betalib

import (
	"os"

	apexlog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Early Logger until Zap is fully configured and bootstrapped. Once bootstrapped, we will forward these logs to Zap
func init() {
	apexlog.SetHandler(text.New(os.Stderr))
	apexlog.SetLevel(apexlog.DebugLevel)
}
