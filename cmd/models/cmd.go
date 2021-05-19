package models

import "github.com/spf13/cobra"

var modelsCmd = &cobra.Command{
	Use:     "models",
	Short:   "run the gorgeous-admin-server-cli to generate models code",
	Long:    "run the gorgeous-admin-server-cli to generate models code",
	Run:     execute,
	Example: "gorgeous-admin-server-cli models -c ./config.json",
}

func Cmd() *cobra.Command {
	return modelsCmd
}
