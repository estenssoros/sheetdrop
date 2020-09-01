package middle

import (
	"net/http"
	"strings"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/labstack/echo/v4"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

func OktaAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken := c.Request().Header["Authorization"]
			jwt, err := validateAccessToken(accessToken)
			if err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			}
			c.Set(constants.ContextUserName, jwt.Claims["sub"].(string))
			return next(c)
		}
	}
}

func validateAccessToken(accessToken []string) (*jwtverifier.Jwt, error) {
	parts := strings.Split(accessToken[0], " ")
	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "https://dev-862336.okta.com/oauth2/default",
		ClaimsToValidate: map[string]string{"aud": "api://default", "cid": "0oascai94rQkJUCfJ4x6"},
	}
	verifier := jwtVerifierSetup.New()
	return verifier.VerifyIdToken(parts[1])
}
