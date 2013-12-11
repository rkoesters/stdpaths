// +build freebsd linux netbsd openbsd

package stdpaths

import (
	"os/exec"
	"strings"
	"testing"
)

func XdgUserDir(s string) string {
	out, err := exec.Command("xdg-user-dir", s).Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func TestDesktop(t *testing.T) {
	if Desktop() != XdgUserDir("DESKTOP") {
		t.Error(Desktop())
	}
}

func TestDocuments(t *testing.T) {
	if Documents() != XdgUserDir("DOCUMENTS") {
		t.Error(Documents())
	}
}

func TestDownload(t *testing.T) {
	if Download() != XdgUserDir("DOWNLOAD") {
		t.Error(Download())
	}
}

func TestMusic(t *testing.T) {
	if Music() != XdgUserDir("MUSIC") {
		t.Error(Music())
	}
}

func TestPictures(t *testing.T) {
	if Pictures() != XdgUserDir("PICTURES") {
		t.Error(Pictures())
	}
}

func TestPublic(t *testing.T) {
	if Public() != XdgUserDir("PUBLICSHARE") {
		t.Error(Public())
	}
}

func TestTemplates(t *testing.T) {
	if Templates() != XdgUserDir("TEMPLATES") {
		t.Error(Templates())
	}
}

func TestVideos(t *testing.T) {
	if Videos() != XdgUserDir("VIDEOS") {
		t.Error(Videos())
	}
}
