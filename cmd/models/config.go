package models

import (
	"fmt"

	"github.com/spf13/viper"

	"gorgeous-admin-server-cli/global"
)

var (
	_config config
)

type config struct {
	ProjectName string   `mapstructure:"project_name"`
	OutPath     string   `mapstructure:"out_path"`
	IntToInt64  bool     `mapstructure:"int_to_int64"`
	XTime       xTime    `mapstructure:"xtime"`
	DB          dataBase `mapstructure:"db"`
	Tables      []table  `mapstructure:"tables"`
}

type dataBase struct {
	Address  string `mapstructure:"address"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type table struct {
	Name       string `mapstructure:"name"`
	SoftDelete bool   `mapstructure:"soft_delete"`
}

type xTime struct {
	Use     bool   `mapstructure:"use"`
	Package string `mapstructure:"package"`
}

func initConfig() {
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
