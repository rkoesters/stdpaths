package userdirs

import (
	"github.com/rkoesters/stdpaths"
	"path/filepath"
)

func init() {
	dirs.Desktop = filepath.Join(stdpaths.Home(), "Desktop")
	dirs.Documents = filepath.Join(stdpaths.Home(), "Documents")
	dirs.Download = filepath.Join(stdpaths.Home(), "Downloads")
	dirs.Music = filepath.Join(stdpaths.Home(), "Music")
	dirs.Pictures = filepath.Join(stdpaths.Home(), "Pictures")
	dirs.Public = filepath.Join(stdpaths.Home(), "Public")
	dirs.Templates = filepath.Join(stdpaths.Home(), "Templates")
	dirs.Videos = filepath.Join(stdpaths.Home(), "Movies")
}
