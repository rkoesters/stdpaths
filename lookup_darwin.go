package stdpaths

// TODO: Mac paths are hard coded. should think about using system apis
// in the future (go1.3 may allow linking with Objective C).

import (
	"path/filepath"
)

const homeEnvVar = "HOME"

func lookupUserCache() string {
	return filepath.Join(Home(), "Library/Caches")
}

func lookupUserConfig() string {
	return filepath.Join(Home(), "Library/Preferences")
}

func lookupUserData() string {
	return filepath.Join(Home(), "Library/Application Support")
}

func lookupUserRuntime() string {
	return UserCache()
}

func lookupSystemConfig() []string {
	return []string{"/Library/Preferences"}
}

func lookupSystemData() []string {
	return []string{"/Library/Application Support"}
}
