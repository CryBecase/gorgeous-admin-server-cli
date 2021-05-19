package file

import (
	"io/ioutil"

	"github.com/mitchellh/colorstring"

	"gorgeous-admin-server-cli/internal/log"
	"gorgeous-admin-server-cli/pkg/util"
)

type DiskFileTask struct {
	File *DiskFile
	Rule *Rule
}

func NewDiskFileTask(file *DiskFile, rule *Rule) *DiskFileTask {
	return &DiskFileTask{
		File: file,
		Rule: rule,
	}
}

func (dft *DiskFileTask) Exec() error {
	if dft.File.IsDir() {
		// 是目录直接创建即可
		err := util.MakeNotExistDir(dft.File.OutPath(), dft.File.Mode())
		if err != nil {
			return err
		}
	} else {
		if util.IsExist(dft.File.OutPath()) {
			if !dft.Rule.AlwaysNotMatched() && !dft.Rule.AlwaysMatched() {
				log.Warning("\r" + colorstring.Color("[magenta][Warning][reset] "+dft.File.OutPath()+" 已存在，是否覆盖？ ([blue]y[reset]/[blue]n[reset]/[blue]ally[reset]/[blue]alln[reset]) "))
			}
			if !dft.Rule.Matched() {
				// 不能覆盖则跳过
				return nil
			}
		}

		err := ioutil.WriteFile(dft.File.OutPath(), []byte(dft.File.OutData()), dft.File.Mode())
		if err != nil {
			return err
		}
	}
	return nil
}
