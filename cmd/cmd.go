package cmd

import (
	"encoding/base64"

	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"

	"gorgeous-admin-server-cli/cmd/models"
	"gorgeous-admin-server-cli/cmd/run"
	"gorgeous-admin-server-cli/cmd/version"
	"gorgeous-admin-server-cli/global"
	"gorgeous-admin-server-cli/internal/log"
)

var rootCmd = &cobra.Command{
	Use:   "gorgeous-admin-server-cli",
	Short: "gorgeous-admin-server-cli is a golang generator.",
}

func Execute() {
	logo, _ := base64.StdEncoding.DecodeString(encodedLogo)
	rootCmd.SetOut(ansi.NewAnsiStdout())
	rootCmd.PersistentFlags().StringVarP(&global.ConfigPath, "config", "c", "./config.json", "配置文件路径: -c [路径]")
	rootCmd.PersistentFlags().BoolVarP(&global.Debug, "debug", "d", false, "是否开启 debug 模式：-d")

	// help
	rootCmd.SetHelpTemplate(colorstring.Color("[cyan]"+string(logo)+"[reset]\n") + "Version:" + global.Version + "\n\n" + rootCmd.HelpTemplate())

	// run
	rootCmd.AddCommand(run.Cmd())

	// version
	rootCmd.AddCommand(version.Cmd())

	// models
	rootCmd.AddCommand(models.Cmd())

	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		return
	}
}

var encodedLogo = "ICAgX19fXyAgICAgICAgICAgICAgICAgXyAgICAgICANCiAgLyBfX198IF9fXyAgXyAgIF8gIF9ffCB8IF9fIF8gDQogfCB8ICBfIC8gXyBcfCB8IHwgfC8gX2AgfC8gX2AgfA0KIHwgfF98IHwgKF8pIHwgfF98IHwgKF98IHwgKF98IHwNCiAgXF9fX198XF9fXy8gXF9fLF98XF9fLF98XF9fLF98DQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIA=="
