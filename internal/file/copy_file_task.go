package file

import (
	"io/ioutil"
	"strings"

	"github.com/markbates/pkger"
	"github.com/mitchellh/colorstring"

	"gorgeous-admin-server-cli/global"
	"gorgeous-admin-server-cli/internal/log"
	"gorgeous-admin-server-cli/pkg/util"
)

// CopyFileTask copy 到磁盘的文件
type CopyFileTask struct {
	File *CopyFile
	Rule *Rule
}

func NewCopyFileTask(file *CopyFile, rule *Rule) *CopyFileTask {
	return &CopyFileTask{
		File: file,
		Rule: rule,
	}
}

func (cft *CopyFileTask) Exec() error {
	if cft.File.IsDir() {
		return util.MakeNotExistDir(cft.File.OutPath(), cft.File.Mode())
	}

	if util.IsExist(cft.File.OutPath()) {
		if !cft.Rule.AlwaysNotMatched() && !cft.Rule.AlwaysMatched() {
			log.Warning("\r" + colorstring.Color("[magenta][Warning][reset] "+cft.File.OutPath()+" 已存在，是否覆盖？ ([blue]y[reset]/[blue]n[reset]/[blue]ally[reset]/[blue]alln[reset]) "))
		}
		if !cft.Rule.Matched() {
			// 不能覆盖则跳过
			return nil
		}
	}

	f, err := pkger.Open(cft.File.SrcPath())
	if err != nil {
		return err
	}
	defer f.Close()

	sourceFile, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	fstr := string(sourceFile)
	fstr = strings.ReplaceAll(fstr, global.Domain, cft.File.ProjectName())

	err = ioutil.WriteFile(cft.File.OutPath(), []byte(fstr), cft.File.Mode())
	if err != nil {
		return err
	}

	return nil
}
