package middle

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/estenssoros/sheetdrop/constants"
	"github.com/labstack/echo/v4"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/pkg/errors"
)

func Auth() echo.MiddlewareFunc {
	return Auth0()
}

func OktaAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken, err := extractAuthToken(c)
			if err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			}
			jwtToken, err := validateOktaAccessToken(accessToken)
			if err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			}
			c.Set(constants.ContextUserName, jwtToken.Claims["sub"].(string))
			return next(c)
		}
	}
}

func validateOktaAccessToken(accessToken string) (*jwtverifier.Jwt, error) {
	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "https://dev-862336.okta.com/oauth2/default",
		ClaimsToValidate: map[string]string{"aud": "api://default", "cid": "0oascai94rQkJUCfJ4x6"},
	}
	verifier := jwtVerifierSetup.New()
	return verifier.VerifyIdToken(accessToken)
}

func Auth0() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken, err := extractAuthToken(c)
			if err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			}
			token, err := validateAuth0Token(accessToken)
			if err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			}
			c.Set(constants.ContextUserName, token.Claims.(jwt.MapClaims)["sub"].(string))
			return next(c)
		}
	}
}

func extractAuthToken(c echo.Context) (string, error) {
	accessToken := c.Request().Header["Authorization"]
	if len(accessToken) == 0 {
		return "", errors.New("no auth token")
	}
	parts := strings.Split(accessToken[0], " ")
	if len(parts) == 0 {
		return "", errors.New("malformed token")
	}
	return parts[1], nil
}

func validateAuth0Token(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		aud := "YOUR_API_IDENTIFIER"
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		iss := "https://dev-km0fs7uh.us.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			return nil, errors.Wrap(err, "getPemCert")
		}

		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	})
	return token, errors.Wrap(err, "jwt.ParseWithClaims")
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://dev-km0fs7uh.us.auth0.com/.well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, errors.Wrap(err, "json.Decode")
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
