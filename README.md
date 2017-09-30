utils
=====

Short, tested, and actively used utilities for Go.

File operations `utils/fileutil`
--------------------------------
```bash
go get github.com/ef2k/utils/fileutil
```

- `IsDir(string) (bool, error)`
- `Exists(string) (bool, error)`
- `TempDir() *fileutil.TmpDir`
- `TempFile() *fileutil.TmpFile`

License
-------
MIT
