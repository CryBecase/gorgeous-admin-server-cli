package run

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gorgeous-admin-server-cli/global"
)

const _runPath = "/cmd/run/kf6vca0s"

var _config = config{
	ProjectName: "awesome-project",
	OutPath:     "./",
}

type config struct {
	ProjectName string `mapstructure:"project_name"`
	OutPath     string `mapstructure:"out_path"`
}

func initConfig(cmd *cobra.Command) {
	p := _config.ProjectName
	o := _config.OutPath

	if cmd.Parent().Flag("config").Changed {
		v := viper.New()
		v.SetConfigFile(global.ConfigPath)
		v.SetConfigType("json")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Open config file error: %s \n", err))
		}

		if err := v.Unmarshal(&_config); err != nil {
			panic(err)
		}
	}

	if cmd.Flag("project name").Changed {
		_config.ProjectName = p
	}
	if cmd.Flag("out path").Changed {
		_config.OutPath = o
	}

	_config.OutPath = strings.Trim(_config.OutPath, " ")
	if _config.OutPath[len(_config.OutPath)-1] != '/' {
		_config.OutPath = _config.OutPath + "/"
	}
}
