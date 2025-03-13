package main

import (
	"runtime"

	cmd "apps/clients/public/cli/v2alpha/oeco/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
