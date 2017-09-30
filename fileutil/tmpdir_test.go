package fileutil

import "testing"

func TestTempDir(t *testing.T) {
	tmpDir, err := TempDir()
	if err != nil {
		t.Fatal(err)
	}
	defer tmpDir.Remove()

	t.Run("check existence", func(t *testing.T) {
		if isDir, _ := IsDir(tmpDir.Path); !isDir {
			t.Fatal("The tmp path should exist")
		}
	})

	t.Run("remove the tempdir", func(t *testing.T) {
		tmpDir.Remove()
		_, err := IsDir(tmpDir.Path)
		if err == nil {
			t.Fatalf("The tempdir should not exist")
		}
	})
}
