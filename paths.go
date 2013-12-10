package stdpaths

import (
	"os"
)

// Home returns the user's home directory.
func Home() string {
	if cache.Home == "" {
		cache.Home = lookupHome()
	}
	return cache.Home
}

// Tmp returns the directory that should be used for
// temporary files.
func Tmp() string {
	if cache.Tmp == "" {
		cache.Tmp = os.TempDir()
	}
	return cache.Tmp
}

// UserCache returns the path to the directory where
// non-essential (cached) data should be written.
func UserCache() string {
	if cache.UserCache == "" {
		cache.UserCache = lookupUserCache()
	}
	return cache.UserCache
}

// UserConfig returns the path to the directory where
// user configuration files should be written.
func UserConfig() string {
	if cache.UserConfig == "" {
		cache.UserConfig = lookupUserConfig()
	}
	return cache.UserConfig
}

// UserData returns the path to the directory where
// user data files should be written.
func UserData() string {
	if cache.UserData == "" {
		cache.UserData = lookupUserData()
	}
	return cache.UserData
}

// UserRuntime returns the path to the directory where
// runtime files should be placed.
func UserRuntime() string {
	if cache.UserRuntime == "" {
		cache.UserRuntime = lookupUserRuntime()
	}
	return cache.UserRuntime
}

// SystemConfig returns a slice of paths that should be
// searched for configuration files.
func SystemConfig() []string {
	if cache.SystemConfig == nil {
		cache.SystemConfig = lookupSystemConfig()
	}
	return cache.SystemConfig
}

// SystemData returns a slice of paths that should be
// searched for data files.
func SystemData() []string {
	if cache.SystemData == nil {
		cache.SystemData = lookupSystemData()
	}
	return cache.SystemData
}
