package userdirs

import (
	"github.com/rkoesters/stdpaths"
	"path/filepath"
)

func initPersonalDirs() {
	if cache != nil {
		return
	}
	cache = new(userDirs)

	cache.Desktop = filepath.Join(stdpaths.Home(), "Desktop")
	cache.Documents = filepath.Join(stdpaths.Home(), "Documents")
	cache.Download = filepath.Join(stdpaths.Home(), "Downloads")
	cache.Music = filepath.Join(stdpaths.Home(), "Music")
	cache.Pictures = filepath.Join(stdpaths.Home(), "Pictures")
	cache.Public = filepath.Join(stdpaths.Home(), "Public")
	cache.Templates = filepath.Join(stdpaths.Home(), "Templates")
	cache.Videos = filepath.Join(stdpaths.Home(), "Movies")
}
