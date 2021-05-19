package run

import "github.com/spf13/cobra"

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "run the gorgeous-admin-server-cli to generate project code",
	Long:    "run the gorgeous-admin-server-cli to generate project code",
	Run:     execute,
	Example: "gorgeous-admin-server-cli run -p awesome-project -o ./",
}

func Cmd() *cobra.Command {
	runCmd.PersistentFlags().StringVarP(&_config.ProjectName, "project name", "p", "awesome-project", "项目名: -p awesome-project")
	runCmd.PersistentFlags().StringVarP(&_config.OutPath, "out path", "o", "./", "输出目录: -o ./")
	return runCmd
}
