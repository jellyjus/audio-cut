package routing

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/login", login)
	e.GET("/set_access_token", setAccessToken)

	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(secret),
		TokenLookup: fmt.Sprintf("cookie:%s", tokenCookie),
		Claims:      &JWTClaims{},
	}))

	r.GET("/get_audios", getAudios)
	r.POST("/upload_audio", uploadAudio)
}

func getUser(c echo.Context) *JWTClaims {
	user := c.Get("user")
	if user == nil {
		return nil
	}
	return user.(*jwt.Token).Claims.(*JWTClaims)
}
