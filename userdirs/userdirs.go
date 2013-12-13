package userdirs

type userDirs struct {
	Desktop   string
	Documents string
	Download  string
	Music     string
	Pictures  string
	Public    string
	Templates string
	Videos    string
}

var dirs = new(userDirs)

// Desktop returns the user's desktop directory.
func Desktop() string {
	return dirs.Desktop
}

// Documents returns the user's documents directory.
func Documents() string {
	return dirs.Documents
}

// Download returns the user's download directory.
func Download() string {
	return dirs.Download
}

// Music returns the user's music directory.
func Music() string {
	return dirs.Music
}

// Pictures returns the user's pictures directory.
func Pictures() string {
	return dirs.Pictures
}

// Public returns the user's public directory.
func Public() string {
	return dirs.Public
}

// Templates returns the user's template directory.
func Templates() string {
	return dirs.Templates
}

// Videos returns the user's videos directory.
func Videos() string {
	return dirs.Videos
}
