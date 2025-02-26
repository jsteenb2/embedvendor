package embedvendor

import (
	"embed"
	"io/fs"
)

//go:embed testdata
var testdataFS embed.FS

func WalkFS() []string {
	var out []string
	err := fs.WalkDir(testdataFS, "testdata", func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		out = append(out, path)
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
	return out
}
