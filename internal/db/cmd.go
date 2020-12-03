package db

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(migrateCmd)
	Cmd.AddCommand(connectCmd)
	Cmd.AddCommand(pingCmd)
}

// Cmd entrypoint
var Cmd = &cobra.Command{
	Use:   "db",
	Short: "",
}
