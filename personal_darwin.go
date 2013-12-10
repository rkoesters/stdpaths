package paths

import (
	"path/filepath"
)

func initPersonalDirs() {
	cache.Desktop = filepath.Join(Home(), "Desktop")
	cache.Documents = filepath.Join(Home(), "Documents")
	cache.Download = filepath.Join(Home(), "Downloads")
	cache.Music = filepath.Join(Home(), "Music")
	cache.Pictures = filepath.Join(Home(), "Pictures")
	cache.Public = filepath.Join(Home(), "Public")
	cache.Templates = filepath.Join(Home(), "Templates")
	cache.Videos = filepath.Join(Home(), "Movies")
}
