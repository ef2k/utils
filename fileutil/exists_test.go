package fileutil

import "testing"

func TestExists(t *testing.T) {
	tempDir, err := TempDir()
	if err != nil {
		t.Fatal(err)
	}
	defer tempDir.Remove()

	t.Run("check for existence", func(t *testing.T) {
		exists, err := Exists(tempDir.Path)
		if err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Fatal("Expected the file to exists")
		}
	})

	t.Run("check non-existence", func(t *testing.T) {
		tempDir.Remove()
		exists, err := Exists(tempDir.Path)
		if err != nil {
			t.Fatal(err)
		}
		if exists {
			t.Fatal("Expected the file to not exist")
		}
	})
}
