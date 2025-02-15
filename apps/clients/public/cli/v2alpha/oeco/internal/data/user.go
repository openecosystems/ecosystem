package data

import (
	"os"
	"os/user"
	"runtime"
)

// GetUserName get the current username for the logged in user
func GetUserName() string {
	username := ""

	if runtime.GOOS == "windows" {
		username = os.Getenv("USERNAME")
	} else {
		current, err := user.Current()
		if err != nil {
			return "oeco"
		}
		username = current.Username
	}

	return username
}
