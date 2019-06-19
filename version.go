package version

type Version struct {
	AppName   string
	Version   string
	BuildTime string
}

func (v *Version) ShowVersion() string {
	return v.AppName + v.Version + v.BuildTime
}
