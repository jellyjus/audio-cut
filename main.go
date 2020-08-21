package main

import (
	"github.com/jellyjus/audio-cut/config"
	"github.com/jellyjus/audio-cut/routing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("ENV")
	if env == "" {
		os.Setenv("ENV", "local")
		env = os.Getenv("ENV")
	}

	if err := config.ParseConfig(env); err != nil {
		panic(err)
	}

	e := echo.New()
	//e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	routing.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
