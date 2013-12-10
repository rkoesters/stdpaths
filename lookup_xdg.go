// +build freebsd linux netbsd openbsd

package paths

import (
	"os"
	"path/filepath"
)

const homeEnvVar = "HOME"

func lookupUserDir(env, dflt string) string {
	dir := os.Getenv(env)
	if len(dir) != 0 {
		return dir
	}
	return filepath.Join(Home(), dflt)
}

func lookupUserCache() string {
	return lookupUserDir("XDG_CACHE_HOME", ".cache")
}

func lookupUserConfig() string {
	return lookupUserDir("XDG_CONFIG_HOME", ".config")
}

func lookupUserData() string {
	return lookupUserDir("XDG_DATA_HOME", ".local/share")
}

func lookupUserRuntime() string {
	runtime := os.Getenv("XDG_RUNTIME_DIR")
	if len(runtime) != 0 {
		return runtime
	}
	return UserCache()
}

func lookupSystemDirs(env, dflt string) []string {
	dirs := os.Getenv(env)
	if len(dirs) == 0 {
		dirs = dflt
	}
	return filepath.SplitList(dirs)
}

func lookupSystemConfig() []string {
	return lookupSystemDirs("XDG_CONFIG_DIRS", "/etc/xdg")
}

func lookupSystemData() []string {
	return lookupSystemDirs("XDG_DATA_DIRS", "/usr/local/share/:/usr/share/")
}
