package routing

import (
	"github.com/jellyjus/audio-cut/vk"
	"github.com/labstack/echo/v4"
	"net/http"
)

type uploadAudioBody struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func uploadAudio(c echo.Context) error {
	user := getUser(c)
	a := new(uploadAudioBody)
	if err := c.Bind(a); err != nil {
		return err
	}

	file, _ := c.FormFile("file")
	fileSrc, err := file.Open()
	defer fileSrc.Close()

	server, err := vk.GetUploadServer(user.AccessToken)
	if err != nil {
		return err
	}

	err = vk.UploadAudio(user.AccessToken, server, fileSrc, file.Filename, a.Artist, a.Title)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "ok")
}
