package cmd

import (
	"github.com/estenssoros/sheetdrop/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(server.Cmd)
}

var cmd = &cobra.Command{
	Use:   "sheetdrop",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}
