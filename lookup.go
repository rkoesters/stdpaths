package stdpaths

import (
	"os"
	"os/user"
	"path/filepath"
)

type lookupCache struct {
	Home string
	Tmp  string

	UserCache   string
	UserConfig  string
	UserData    string
	UserRuntime string

	SystemConfig []string
	SystemData   []string
}

var cache = new(lookupCache)

func lookupHome() string {
	// Let's try the OS specific way first.
	home := os.Getenv(homeEnvVar)
	if len(home) != 0 {
		return home
	}

	// Didn't get it? Let's try using the standard library.
	u, err := user.Current()
	if err == nil {
		return u.HomeDir
	}

	// Crap. Now we fallback on using the temp dir.
	return filepath.Join(Tmp(), username())
}
