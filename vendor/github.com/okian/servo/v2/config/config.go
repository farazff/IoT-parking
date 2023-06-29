package config

var (
	// The values must be set on build time
	app     string
	version string
	commit  string
	tag     string
	branch  string
	// date of build
	date string
)

func AppName() string {
	return app
}

func Version() string {
	return version
}

func Commit() string {
	return commit
}

func Tag() string {
	return tag
}

func Branch() string {
	return branch
}

func Date() string {
	return date
}
