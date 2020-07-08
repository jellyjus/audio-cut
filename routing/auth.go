package routing

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jellyjus/audio-cut/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

const authRedirectURL = "https://oauth.vk.com/authorize?client_id=7526254&scope=audio&response_type=code&v=5.118"
const getAccessTokenURL = "https://oauth.vk.com/access_token?client_id=7526254&client_secret=ZSggAUt4jp6rS9Akcmbh"
const tokenCookie = "token"
const secret = "PutinIdet"

type JWTClaims struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserId      int    `json:"user_id"`
	jwt.StandardClaims
}

type redirect struct {
	NeedRedirect bool   `json:"need_redirect"`
	RedirectURL  string `json:"redirect_url"`
}

func login(c echo.Context) error {
	cookie, err := c.Cookie(tokenCookie)
	if err != nil {
		u, _ := url.Parse(authRedirectURL)
		q := u.Query()
		q.Set("redirect_uri", fmt.Sprintf("http://%s%s", c.Request().Host, config.Config.Vk.RedirectURI))
		u.RawQuery = q.Encode()
		return c.JSON(http.StatusOK, redirect{true, u.String()})
	}

	token, _ := jwt.ParseWithClaims(cookie.Value, &JWTClaims{}, nil)
	if token == nil {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
	}

	claims := token.Claims.(*JWTClaims)
	return c.JSON(http.StatusOK, claims)
}

func setAccessToken(c echo.Context) error {
	if c.QueryParam("error") != "" {
		return echo.NewHTTPError(http.StatusInternalServerError, c.QueryParam("error_description"))
	}

	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "code is required")
	}

	u, _ := url.Parse(getAccessTokenURL)
	q := u.Query()
	q.Set("code", code)
	q.Set("redirect_uri", config.Config.Vk.RedirectURI)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	jwtClaims := JWTClaims{}
	err = json.NewDecoder(res.Body).Decode(&jwtClaims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString([]byte(secret))
	cookie := new(http.Cookie)
	cookie.Name = tokenCookie
	cookie.Value = token
	cookie.MaxAge = jwtClaims.ExpiresIn
	c.SetCookie(cookie)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
