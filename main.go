package main

import (
	"github.com/jellyjus/audio-cut/config"
	"github.com/jellyjus/audio-cut/routing"
	"github.com/labstack/echo/v4"
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
	routing.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
