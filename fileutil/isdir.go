package fileutil

import (
	"os"
)

// Checks if the given path is a directory. Returns an error only if existence
// can't be confirmed or something unrelated goes wrong.
func IsDir(path string) (bool, error) {

	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return f.Mode().IsDir(), nil
}
