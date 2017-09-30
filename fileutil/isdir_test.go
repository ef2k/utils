package fileutil

import "testing"

func TestIsDir(t *testing.T) {
	t.Run("confirm directory", func(t *testing.T) {
		tmpdir, err := TempDir()
		if err != nil {
			t.Fatal(err)
		}
		defer tmpdir.Remove()

		if isDir, err := IsDir(tmpdir.Path); !isDir || err != nil {
			t.Fatal("Expected the path to be confirmed a directory")
		}
	})

	t.Run("confirm not directory", func(t *testing.T) {
		tmpf, err := TempFile()
		if err != nil {
			t.Fatal(err)
		}
		defer tmpf.Remove()

		if isDir, err := IsDir(tmpf.Path); isDir || err != nil {
			t.Fatal("Expected the path to not be a directory")
		}
	})
}
