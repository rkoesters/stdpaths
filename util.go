package stdpaths

import (
	"os"
	"os/user"
)

// username tries its very best to get the username.
func username() string {
	u, err := user.Current()
	if err == nil {
		return u.Username
	}

	// We couldn't get the username, so lets just use
	// the name of the executable.
	return os.Args[0]
}
