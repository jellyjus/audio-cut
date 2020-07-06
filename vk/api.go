package vk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

const baseURL = "https://api.vk.com/method"

type response struct {
	Response struct {
		Count int     `json:"count"`
		Items []Audio `json:"items"`
	} `json:"response"`
}

type Audio struct {
	Id      int    `json:"id"`
	OwnerId int    `json:"owner_id"`
	Artist  string `json:"artist"`
	Title   string `json:"title"`
	Url     string `json:"url"`
}

var client = &http.Client{}

func GetAudios(accessToken string, userId int) ([]Audio, error) {
	v := url.Values{}
	v.Set("owner_id", strconv.Itoa(userId))
	v.Set("count", "10")
	v.Set("access_token", accessToken)
	v.Set("v", "5.91")

	uri := fmt.Sprintf("%s/audio.get?%s", baseURL, v.Encode())
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", "VKAndroidApp/5.52-4543 (Android 5.1.1; SDK 22; x86_64; unknown Android SDK built for x86_64; en; 320x240)")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var data response
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data.Response.Items, nil
}

func GetUploadServer(accessToken string) (string, error) {
	v := url.Values{}
	v.Set("access_token", accessToken)
	v.Set("v", "5.120")

	uri := fmt.Sprintf("%s/audio.getUploadServer?%s", baseURL, v.Encode())
	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	body := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&body)
	return body["response"].(map[string]interface{})["upload_url"].(string), nil
}

func UploadAudio(accessToken string, server string, file io.Reader, filename string, artist string, title string) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part1, err := writer.CreateFormFile("file", filename)
	_, err = io.Copy(part1, file)
	if err != nil {
		return err
	}
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", server, payload)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	defer res.Body.Close()
	resUpload := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&resUpload)

	v := url.Values{}
	v.Set("access_token", accessToken)
	v.Set("server", fmt.Sprintf("%.0f", resUpload["server"].(float64)))
	v.Set("audio", resUpload["audio"].(string))
	v.Set("hash", resUpload["hash"].(string))
	v.Set("artist", artist)
	v.Set("title", title)
	v.Set("v", "5.120")

	uri := fmt.Sprintf("%s/audio.save?%s", baseURL, v.Encode())
	res, err = http.Get(uri)
	if err != nil {
		return err
	}
	resp, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resp))

	return nil
}
