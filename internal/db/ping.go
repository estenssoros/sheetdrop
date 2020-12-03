package db

import (
	"github.com/estenssoros/sheetdrop/orm"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:     "ping",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := orm.Connect()
		if err != nil {
			return errors.Wrap(err, "orm.Connect")
		}
		return nil
	},
}
