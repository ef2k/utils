package fileutil

import (
	"io/ioutil"
	"os"
)

func TempDir() (*TmpDir, error) {
	p, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}
	return &TmpDir{p}, nil
}

type TmpDir struct {
	Path string
}

func (t *TmpDir) Remove() error {
	return os.RemoveAll(t.Path)
}
