package stdpaths

import (
	"path/filepath"
)

const homeEnvVar = "HOME"

func lookupUserCache() string {
	return filepath.Join(Home(), "Library", "Caches")
}

func lookupUserConfig() string {
	return filepath.Join(Home(), "Library", "Preferences")
}

func lookupUserData() string {
	return filepath.Join(Home(), "Library", "Application Support")
}

func lookupUserRuntime() string {
	return UserCache()
}

func lookupSystemConfig() []string {
	// TODO: I am not sure if this is correct
	return []string{}
}

func lookupSystemData() []string {
	// TODO: I am not sure if this is correct
	return []string{}
}
