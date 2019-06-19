package version

var (
	AppName   string
	Version   string
	BuildTime string
)

func ShowVersion() string {
	return AppName + Version + BuildTime
}
