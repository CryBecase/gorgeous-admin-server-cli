package run

import (
	"fmt"

	"github.com/spf13/cobra"

	"gorgeous-admin-server-cli/internal/file"
	"gorgeous-admin-server-cli/internal/log"
	"gorgeous-admin-server-cli/pkg/task"
)

func execute(cmd *cobra.Command, args []string) {
	log.Debug(fmt.Sprintf("ignored files: %s", file.IgnoredFiles))

	// 初始化配置
	initConfig(cmd)

	// 生成项目文件
	res, err := generate()
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info(res)
}

func generate() (result string, err error) {
	// 获取工作目录下的所有子文件
	files, err := file.GetAllSubFiles(_runPath, _config.OutPath, _config.ProjectName)
	if err != nil {
		log.Debug(fmt.Sprintf("file.GetAllSubFiles error: %s", err.Error()))
	}

	tasks := make([]task.Task, 0)
	rule := file.NewRule()
	for i := range files {
		tasks = append(tasks, file.NewCopyFileTask(&files[i], rule))
	}
	group := task.NewGroup("generating server code...", tasks)
	err = group.Exec()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("项目[[blue]%s[reset]]初始化成功!请到[[blue]%s[reset]]目录下查看!", _config.ProjectName, _config.OutPath+_config.ProjectName), nil
}
