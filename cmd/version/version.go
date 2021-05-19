package version

import (
	"github.com/spf13/cobra"

	"gorgeous-admin-server-cli/global"
	"gorgeous-admin-server-cli/internal/log"
)

func execute(cmd *cobra.Command, args []string) {
	log.Info("Current Version:" + global.Version)
}
