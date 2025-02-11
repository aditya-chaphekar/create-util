package utils

func CheckIfDirExists(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckIfDirEmpty(dir string) bool {
	entries, _ := os.ReadDir(dir)
	if len(entries) == 0 {
		return true
	}
	return false
}
