package seed

import (
	"github.com/estenssoros/sheetdrop/internal/helpers"
	"github.com/estenssoros/sheetdrop/models"
	"github.com/estenssoros/sheetdrop/orm"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "seed",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return seed() },
}

func userSeeds() []interface{} {
	return []interface{}{
		&models.User{
			UserName: "sebastian",
		},
	}
}

func organizationSeeds() []interface{} {
	return []interface{}{
		&models.Organization{
			Name: helpers.StringPtr("estenssoros"),
		},
	}
}

func organizationUserSeeds() []interface{} {
	return []interface{}{
		&models.OrganizationUser{
			OrganizationID: 1,
			UserID:         1,
		},
	}
}

func resourceSeeds() []interface{} {
	return []interface{}{
		&models.Resource{
			OrganizationID: 1,
			OwnerID:        1,
			Name:           helpers.StringPtr("myResource"),
		},
	}
}

func schemaSeeds() []interface{} {
	return []interface{}{
		&models.Schema{
			ResourceID:  1,
			Name:        helpers.StringPtr("mySchema"),
			StartRow:    5,
			StartColumn: 10,
			SourceType:  "excel",
		},
	}
}

func headerSeeds() []interface{} {
	return []interface{}{
		&models.Header{},
	}
}

func headerHeaderSeeds() []interface{} {
	return []interface{}{}
}

func seeds() []interface{} {
	seeds := []interface{}{}
	seeds = append(seeds, userSeeds()...)
	seeds = append(seeds, organizationSeeds()...)
	seeds = append(seeds, organizationUserSeeds()...)
	seeds = append(seeds, resourceSeeds()...)
	seeds = append(seeds, schemaSeeds()...)
	seeds = append(seeds, headerSeeds()...)
	seeds = append(seeds, headerHeaderSeeds()...)
	return seeds
}

func seed() error {
	db, err := orm.Connect()
	if err != nil {
		return errors.Wrap(err, "orm.Connect")
	}
	for _, s := range seeds() {
		if err := db.Create(s).Error; err != nil {
			return errors.Wrap(err, "db.Create")
		}
	}
	return db.Error
}
