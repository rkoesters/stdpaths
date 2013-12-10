// +build freebsd linux netbsd openbsd

package stdpaths

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func initPersonalDirs() {
	if cache.PersonalDirsLoaded {
		return
	}

	f, err := os.Open(filepath.Join(UserConfig(), "user-dirs.dirs"))
	if err != nil {
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := strings.SplitN(strings.TrimSpace(sc.Text()), "=", 2)
		switch l[0] {
		case "XDG_DESKTOP_DIR":
			cache.Desktop = clean(l[1])
		case "XDG_DOCUMENTS_DIR":
			cache.Documents = clean(l[1])
		case "XDG_DOWNLOAD_DIR":
			cache.Download = clean(l[1])
		case "XDG_MUSIC_DIR":
			cache.Music = clean(l[1])
		case "XDG_PICTURES_DIR":
			cache.Pictures = clean(l[1])
		case "XDG_PUBLICSHARE_DIR":
			cache.Public = clean(l[1])
		case "XDG_TEMPLATES_DIR":
			cache.Templates = clean(l[1])
		case "XDG_VIDEOS_DIR":
			cache.Videos = clean(l[1])
		}
	}
	if sc.Err() != nil {
		return
	}
	cache.PersonalDirsLoaded = true
}

func clean(s string) string {
	s = strings.Trim(s, " \t\"")
	if strings.HasPrefix(s, "$HOME") {
		s = filepath.Join(Home(), strings.TrimPrefix(s, "$HOME"))
	}
	return s
}
