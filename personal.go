package paths

// Desktop returns the user's desktop directory.
func Desktop() string {
	initPersonalDirs()
	return cache.Desktop
}

// Documents returns the user's documents directory.
func Documents() string {
	initPersonalDirs()
	return cache.Documents
}

// Download returns the user's download directory.
func Download() string {
	initPersonalDirs()
	return cache.Download
}

// Music returns the user's music directory.
func Music() string {
	initPersonalDirs()
	return cache.Music
}

// Pictures returns the user's pictures directory.
func Pictures() string {
	initPersonalDirs()
	return cache.Pictures
}

// Public returns the user's public directory.
func Public() string {
	initPersonalDirs()
	return cache.Public
}

// Templates returns the user's template directory.
func Templates() string {
	initPersonalDirs()
	return cache.Templates
}

// Videos returns the user's videos directory.
func Videos() string {
	initPersonalDirs()
	return cache.Videos
}
