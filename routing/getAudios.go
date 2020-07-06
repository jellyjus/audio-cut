package routing

import (
	"github.com/jellyjus/audio-cut/config"
	"github.com/jellyjus/audio-cut/vk"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getAudios(c echo.Context) error {
	user := getUser(c)
	audios, err := vk.GetAudios(config.Config.Vk.AccessToken, user.UserId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, audios)
}
