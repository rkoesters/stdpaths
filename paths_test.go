package stdpaths

import (
	"testing"
)

func TestPaths(t *testing.T) {
	t.Log(Home())
	t.Log(SystemConfig())
	t.Log(SystemData())
	t.Log(Tmp())
	t.Log(UserCache())
	t.Log(UserConfig())
	t.Log(UserData())
	t.Log(UserRuntime())
}
