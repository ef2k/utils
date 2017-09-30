package fileutil

import "os"

// Exists checks if the given file path exists. Expect an error only if something
// unexpected goes wrong (permissions, disk failure, etc.) as the error will
// be unrelated to the file's existence.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	return true, nil
}
