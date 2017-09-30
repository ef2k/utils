package fileutil

import (
	"io/ioutil"
	"os"
)

func TempFile() (*TmpFile, error) {
	f, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		return nil, err
	}
	p := f.Name()
	return &TmpFile{p}, nil
}

type TmpFile struct {
	Path string
}

func (t *TmpFile) Remove() error {
	return os.Remove(t.Path)
}
