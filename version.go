package version

import (
	"fmt"
)

var (
	AppName   string
	Version   string
	BuildTime string
)

func ShowVersion() string {
	return fmt.Sprintf("%s %s %s", AppName, Version, BuildTime)
}
