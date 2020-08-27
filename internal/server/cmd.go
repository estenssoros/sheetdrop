package server

import (
	"net/http"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/estenssoros/sheetdrop/controllers"
	"github.com/estenssoros/sheetdrop/internal/middle"
	"github.com/estenssoros/sheetdrop/orm"
	"github.com/estenssoros/sheetdrop/responses"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
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
	t, err := NewTemplate()
	if err != nil {
		return errors.Wrap(err, "NewTemplate")
	}
	e.Renderer = t
	db, err := orm.Connect()
	if err != nil {
		return errors.Wrap(err, "orm.Connect")
	}
	defer db.Close()
	e.Use(middle.DBInjector(db))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/", main, middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	routes(e.Group(
		"/api",
		// middleware.JWT([]byte("secret")),
	))

	return e.Start(":1323")
}

// Main main template
func main(c echo.Context) error {
	sess, _ := session.Get("session", c)
	loggedIn, ok := sess.Values[constants.SessionLoggedIn].(bool)
	if ok && loggedIn {
		user, ok := sess.Values[constants.SessionUser].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, "no session user")
		}
		apis, err := controllers.GetUserAPIs(c.Get(constants.ContextDB).(*gorm.DB), user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.Render(http.StatusOK, "main", &responses.Main{
			LoggedIn: loggedIn,
			APIs:     apis,
		})
	}
	return c.Render(http.StatusOK, "login", c.Get("csrf").(string))
}
