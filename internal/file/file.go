package file

import (
	"os"
	"strings"

	"github.com/markbates/pkger"
)

var (
	IgnoredFiles = []string{".idea"}
)

// GetAllSubFiles 获取path目录下所有的文件(包括子目录)
func GetAllSubFiles(dirPath, outPath, pname string) ([]CopyFile, error) {
	files := make([]CopyFile, 0)

	err := pkger.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, ignoredFile := range IgnoredFiles {
			if strings.Contains(path, ignoredFile) {
				return nil
			}
		}

		pkgerPath := path

		// gorgeous-admin-server-cli:/cmd/run/tmpl/xxx.go => /cmd/run/tmpl/xxx.go
		path = strings.Split(path, ":")[1]
		// /cmd/run/tmpl/xxx.go => /xxx.go
		path = strings.ReplaceAll(path, dirPath, "")
		outPath := outPath + pname + path

		files = append(files, CopyFile{
			srcPath: pkgerPath,
			pname:   pname,
			DiskFile: DiskFile{
				info:    info,
				outPath: outPath,
			},
		})

		return nil
	})

	return files, err
}
