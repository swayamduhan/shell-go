package utils

import "os"

func HasBash() (string, bool) {
	possiblePaths := []string{
		`C:\Program Files\Git\bin\bash.exe`,
		`C:\Program Files (x86)\Git\bin\bash.exe`,
		`C:\Git\bin\bash.exe`,
	}
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path, true
		}
	}
	return "", false
}