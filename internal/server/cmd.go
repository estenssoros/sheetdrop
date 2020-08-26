package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/internal/middle"
	"github.com/estenssoros/sheetdrop/orm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	t, err := NewTemplate()
	if err != nil {
		return errors.Wrap(err, "NewTemplate")
	}
	e.Renderer = t
	e.GET("/", Main)
	db, err := orm.Connect()
	if err != nil {
		return errors.Wrap(err, "orm.Connect")
	}
	defer db.Close()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	routes(e.Group(
		"/api",
		// middleware.JWT([]byte("secret")),
		middle.DBInjector(db),
	))

	return e.Start(":1323")
}

// Main main template
func Main(c echo.Context) error {
	return c.Render(http.StatusOK, "main", "World")
}
