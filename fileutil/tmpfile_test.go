package fileutil

import "testing"

func TestTempFile(t *testing.T) {
	f, err := TempFile()
	if err != nil {
		t.Fatal(err)
	}
	defer f.Remove()

	t.Run("check existence", func(t *testing.T) {
		exists, err := Exists(f.Path)
		if err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Error("Expected the file to exist")
		}
	})

	t.Run("should not exist", func(t *testing.T) {
		f.Remove()
		exists, err := Exists(f.Path)
		if err != nil {
			t.Fatal(err)
		}
		if exists {
			t.Error("Expected the file not to exist")
		}
	})
}
