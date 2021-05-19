package version

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "show gorgeous-admin-server-cli version",
	Long:    "show gorgeous-admin-server-cli version",
	Run:     execute,
	Example: "gorgeous-admin-server-cli version",
}

func Cmd() *cobra.Command {
	return versionCmd
}
