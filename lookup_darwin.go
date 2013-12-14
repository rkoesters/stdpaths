package stdpaths

// TODO: Mac paths are hard coded. should think about using system apis
// in the future (go1.3 may allow linking with Objective C).

import (
	"path/filepath"
)

const (
	homeEnvVar = "HOME"

	macCache  = "/Library/Caches"
	macConfig = "/Library/Preferences"
	macData   = "/Library/Application Support"
)

func lookupUserCache() string {
	return filepath.Join(Home(), macCache)
}

func lookupUserConfig() string {
	return filepath.Join(Home(), macConfig)
}

func lookupUserData() string {
	return filepath.Join(Home(), macData)
}

func lookupUserRuntime() string {
	return UserCache()
}

func lookupSystemConfig() []string {
	return []string{macConfig}
}

func lookupSystemData() []string {
	return []string{macData}
}
