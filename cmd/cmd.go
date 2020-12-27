package cmd

import (
	"github.com/estenssoros/sheetdrop/internal/db"
	"github.com/estenssoros/sheetdrop/internal/seed"
	"github.com/estenssoros/sheetdrop/server"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(db.Cmd)
	cmd.AddCommand(seed.Cmd)
	cmd.AddCommand(server.Cmd)
}

var cmd = &cobra.Command{
	Use:   "sheetdrop",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}
