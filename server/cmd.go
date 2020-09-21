package server

import (
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/orm"
	"github.com/estenssoros/sheetdrop/server/middle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Cmd entrypoint
var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "run sheetdrop server",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return run() },
}

func run() error {
	e := echo.New()
	db, err := orm.Connect()
	if err != nil {
		return errors.Wrap(err, "orm.Connect")
	}
	e.Use(middle.DBInjector(controllers.New(db)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	restRoutes(e.Group(
		"/rest",
	))
	routes(e.Group(
		"/api",
		middle.Auth(),
	))

	return e.Start(":1323")
}
