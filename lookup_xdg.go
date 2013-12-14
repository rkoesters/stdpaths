// +build freebsd linux netbsd openbsd

package stdpaths

import (
	"os"
	"path/filepath"
)

const homeEnvVar = "HOME"

func lookupUserCache() string {
	dir := os.Getenv("XDG_CACHE_HOME")
	if len(dir) != 0 {
		return dir
	}
	return filepath.Join(Home(), ".cache")
}

func lookupUserConfig() string {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if len(dir) != 0 {
		return dir
	}
	return filepath.Join(Home(), ".config")
}

func lookupUserData() string {
	dir := os.Getenv("XDG_DATA_HOME")
	if len(dir) != 0 {
		return dir
	}
	return filepath.Join(Home(), ".local/share")
}

func lookupUserRuntime() string {
	runtime := os.Getenv("XDG_RUNTIME_DIR")
	if len(runtime) != 0 {
		return runtime
	}
	return UserCache()
}

func lookupSystemConfig() []string {
	dirs := os.Getenv("XDG_CONFIG_DIRS")
	if len(dirs) != 0 {
		return filepath.SplitList(dirs)
	}
	return []string{"/etc/xdg"}
}

func lookupSystemData() []string {
	dirs := os.Getenv("XDG_DATA_DIRS")
	if len(dirs) != 0 {
		return filepath.SplitList(dirs)
	}
	return []string{"/usr/local/share", "/usr/share"}
}
