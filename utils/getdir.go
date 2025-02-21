package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home != "" {
		return filepath.Clean(home), nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Clean(usr.HomeDir), nil
}

func GetDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = filepath.Clean(dir)

	home, err := GetHomeDir()
	if err != nil {
		return dir
	}

	if dir == home {
		return ""
	}

	if strings.HasPrefix(dir, home+string(os.PathSeparator)) {
		dir = "~" + strings.TrimPrefix(dir, home)
	}

	return dir
}