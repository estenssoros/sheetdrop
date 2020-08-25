package db

import (
	"os"
	"os/exec"
	"strings"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Cmd entrypoint
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to database",
	RunE: func(cmd *cobra.Command, args []string) error {
		cArgs := []string{
			"-u" +
				os.Getenv(constants.EnvDataBaseUser),
			"-h" +
				os.Getenv(constants.EnvDataBaseHost),
			"-p" +
				os.Getenv(constants.EnvDataBasePassword),
			os.Getenv(constants.EnvDataBaseName),
			"-A",
		}
		if len(args) > 1 {
			cArgs = append(cArgs, args[0])
		}
		color.Green("mysql " + strings.Join(cArgs, " "))
		c := exec.Command("mysql", cArgs...)
		{
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Stdout = os.Stdout
		}
		if err := c.Run(); err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}
