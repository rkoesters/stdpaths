package paths

import (
	"testing"
)

func TestPersonal(t *testing.T) {
	t.Log(Desktop())
	t.Log(Documents())
	t.Log(Download())
	t.Log(Music())
	t.Log(Pictures())
	t.Log(Public())
	t.Log(Templates())
	t.Log(Videos())
}
