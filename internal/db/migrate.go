package db

import (
	"github.com/estenssoros/sheetdrop/internal/models"
	"github.com/estenssoros/sheetdrop/orm"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var drop bool

func init() {
	migrateCmd.Flags().BoolVarP(&drop, "drop", "d", false, "drop models during migration")
}

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "migrates models to database",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := orm.Connect()
		if err != nil {
			return errors.Wrap(err, "orm.Connect")
		}
		for _, model := range models.Models {
			if drop {
				if err := db.Migrator().DropTable(model); err != nil {
					return errors.Wrap(err, "orm.DropTableMsSQL")
				}
			}
			if err := db.AutoMigrate(model); err != nil {
				return errors.Wrap(err, "db.Automigrate")
			}
		}
		return nil
	},
}
