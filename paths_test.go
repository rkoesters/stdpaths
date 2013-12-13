package stdpaths

import (
	"testing"
)

func TestPaths(t *testing.T) {
	// TODO: add real automatic tests.
	t.Log(Home())
	t.Log(SystemConfig())
	t.Log(SystemData())
	t.Log(Tmp())
	t.Log(UserCache())
	t.Log(UserConfig())
	t.Log(UserData())
	t.Log(UserRuntime())
}
